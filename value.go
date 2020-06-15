package gohaystack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
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
		/*
			case `n`:
			case `d`:
			case `h`:
			case `t`:
			case `u`:
			case `c`:
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
