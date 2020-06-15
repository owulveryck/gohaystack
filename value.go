package gohaystack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Kind is a supported type for a haystack value
type Kind int

// Unit represents a unit is a number
type Unit *string

// NewUnit returns a new unit
func NewUnit(u string) Unit {
	return &u
}

// NewNumber ...
func NewNumber(value float32, unit Unit) *Value {
	return &Value{
		kind: HaystackTypeNumber,
		number: struct {
			value float32
			unit  Unit
		}{
			value,
			unit,
		},
	}
}

// NewURL ...
func NewURL(u *url.URL) *Value {
	return &Value{
		kind: HaystackTypeURI,
		u:    u,
	}

}

// GetKind of the underlying value
func (v *Value) GetKind() Kind {
	return v.kind
}

// Value is an haystack value
type Value struct {
	kind   Kind
	str    *string
	number struct {
		value float32
		unit  Unit
	}
	b     bool
	t     time.Time
	u     *url.URL
	ref   *HaystackID
	g     *Grid
	dict  map[string]*Value
	list  []*Value
	coord struct {
		long float32
		lat  float32
	}
}

func (v *Value) unmarshalJSONNotString(b []byte) error {
	var list []*Value
	err := json.Unmarshal(b, &list)
	if err == nil {
		v.kind = HaystackTypeList
		v.list = list
		return nil
	}
	var dict map[string]*Value
	err = json.Unmarshal(b, &dict)
	if err == nil {
		v.kind = HaystackTypeDict
		v.dict = dict
		return nil
	}
	var boolean bool
	err = json.Unmarshal(b, &boolean)
	if err == nil {
		v.kind = HaystackTypeBool
		v.b = boolean
		return nil
	}
	return errors.New("not implemented")
}

func (v *Value) unmarshalJSONString(b []byte) error {
	re := regexp.MustCompile(`^(([m\-znrsdhtucx]):)?(.*)$`)
	res := re.FindStringSubmatch(string(b))
	switch res[2] {
	case ``:
		v.kind = HaystackTypeStr
		val := res[3]
		v.str = &val
	case `m`:
		v.kind = HaystackTypeMarker
	case `-`:
		v.kind = HaystackTypeRemove
	case `z`:
		v.kind = HaystackTypeNA
	case "n":
		elements := strings.Fields(res[3])
		f, err := strconv.ParseFloat(elements[0], 32)
		if err != nil {
			return err
		}
		var unit string
		if len(elements) > 2 {
			return errors.New("wrong entry for number, too many elements")
		}
		if len(elements) > 1 {
			unit = elements[1]
			v.number.unit = &unit
		}
		v.kind = HaystackTypeNumber
		v.number.value = float32(f)
	case `d`:
		t, err := time.Parse("2006-01-02", res[3])
		if err != nil {
			return err
		}
		v.kind = HaystackTypeDate
		v.t = t
	case `h`:
		t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "1970-01-01 "+res[3]+" +0000 UTC")
		if err != nil {
			return err
		}
		v.kind = HaystackTypeTime
		v.t = t
	case `t`:
		// TODO: handle extra location
		t, err := time.Parse(time.RFC3339, res[3])
		if err != nil {
			return err
		}
		v.kind = HaystackTypeDateTime
		v.t = t
	case `u`:
		u, err := url.Parse(res[3])
		if err != nil {
			return err
		}
		v.kind = HaystackTypeURI
		v.u = u
	case `c`:
		elements := strings.Split(res[3], ",")
		if len(elements) != 2 {
			return errors.New("bad coordinates, expected lat,long")
		}
		lat, err := strconv.ParseFloat(elements[0], 32)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(elements[1], 32)
		if err != nil {
			return err
		}
		v.kind = HaystackTypeCoord
		v.coord.lat = float32(lat)
		v.coord.long = float32(long)
		/*
			case `x`:
		*/
	case `s`:
		v.kind = HaystackTypeStr
		val := res[3]
		v.str = &val
	case `r`:
		v.kind = HaystackTypeRef
		val := res[3]
		id := HaystackID(val)
		v.ref = &id
	default:
		return errors.New("not implemented")
	}
	return nil
}

// UnmarshalJSON extract a value from b
func (v *Value) UnmarshalJSON(b []byte) error {
	if b == nil || len(b) == 0 {
		return errors.New("Cannot unmarshal nil or empty value")
	}
	// is it a string
	if isValidString(b) {
		return v.unmarshalJSONString(trimDoubleQuote(b))
	}
	return v.unmarshalJSONNotString(b)
}

// MarshalJSON encode the value in format compatible with haystack's JSON:
// https://www.project-haystack.org/doc/Json
func (v *Value) MarshalJSON() ([]byte, error) {
	var output string
	switch v.kind {
	case HaystackTypeBool:
		output = fmt.Sprintf("%v", v.b)
	case HaystackTypeDict:
		return json.Marshal(v.dict)
	case HaystackTypeList:
		return json.Marshal(v.list)
	case HaystackTypeGrid:
		return json.Marshal(v.g)
	case HaystackTypeStr:
		output = `"s:` + *v.str + `"`
	case HaystackTypeRef:
		output = `"r:` + string(*v.ref) + `"`
	case HaystackTypeRemove:
		output = `"-:"`
	case HaystackTypeMarker:
		output = `"m:"`
	case HaystackTypeNA:
		output = `"z:"`
	case HaystackTypeCoord:
		output = fmt.Sprintf(`"c:%v,%v"`, v.coord.lat, v.coord.long)
	case HaystackTypeDate:
		output = `"d:` + v.t.Format("2006-01-02") + `"`
	case HaystackTypeTime:
		output = `"h:` + v.t.Format("15:04:05") + `"`
	case HaystackTypeDateTime:
		output = `"t:` + v.t.Format(time.RFC3339) + `"`
	case HaystackTypeURI:
		output = `"u:` + (*v.u).String() + `"`
	case HaystackTypeNumber:
		var unit string
		if v.number.unit != nil {
			unit = ` ` + *v.number.unit
		}
		output = `"n:` + strconv.FormatFloat(float64(v.number.value), 'f', -1, 32) + unit + `"`
	default:
		return nil, errors.New("type not handled")
	}
	return []byte(output), nil
}

// NewRef new reference
func NewRef(r *HaystackID) *Value {
	return &Value{
		kind: HaystackTypeRef,
		ref:  r,
	}
}

// NewStr new string value
func NewStr(s string) *Value {
	return &Value{
		kind: HaystackTypeStr,
		str:  &s,
	}
}

// MarkerValue ...
var MarkerValue = &Value{
	kind: HaystackTypeMarker,
}

// GetString value; returns an error if the underlying type is not an haystack string
func (v *Value) GetString() (string, error) {
	if v.kind != HaystackTypeStr {
		return "", errors.New("value type is not a string")
	}
	return *v.str, nil
}

const (
	// HaystackTypeUndefined ...
	HaystackTypeUndefined Kind = iota
	// HaystackTypeGrid is a Grid object
	HaystackTypeGrid
	// HaystackTypeList Array
	HaystackTypeList
	// HaystackTypeDict Object
	HaystackTypeDict
	// HaystackTypeNull null
	HaystackTypeNull
	// HaystackTypeBool Boolean
	HaystackTypeBool
	// HaystackTypeMarker "m:"
	HaystackTypeMarker
	// HaystackTypeRemove "-:"
	HaystackTypeRemove
	// HaystackTypeNA "z:"
	HaystackTypeNA
	// HaystackTypeNumber "n:<float> [unit]" "n:45.5" "n:73.2 °F" "n:-INF"
	HaystackTypeNumber
	// HaystackTypeRef "r:<id> [dis]"  "r:abc-123" "r:abc-123 RTU #3"
	HaystackTypeRef
	// HaystackTypeStr "hello" "s:hello"
	HaystackTypeStr
	// HaystackTypeDate "d:2014-01-03"
	HaystackTypeDate
	// HaystackTypeTime "h:23:59:00"
	HaystackTypeTime
	// HaystackTypeDateTime "t:2015-06-08T15:47:41-04:00 New_York"
	HaystackTypeDateTime
	// HaystackTypeURI "u:http://project-haystack.org/"
	HaystackTypeURI
	// HaystackTypeCoord "c:<lat>,<lng>" "c:37.545,-77.449"
	HaystackTypeCoord
	//HaystackTypeXStr "x:Type:value"
	HaystackTypeXStr
	// HaystackLastType ...
	HaystackLastType
)
