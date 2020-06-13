package gohaystack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type kind int

// Unit represents a unit is a number
type Unit *string

// NewUnit returns a new unit
func NewUnit(u string) Unit {
	return &u
}

// NewNumber ...
func NewNumber(value float32, unit Unit) *Value {
	return &Value{
		kind: haystackTypeNumber,
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
		kind: haystackTypeURI,
		u:    u,
	}

}

// Value is an haystack value
type Value struct {
	kind   kind
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
	coord struct {
		long float32
		lat  float32
	}
}

// UnmarshalJSON extract a value from b
func (v *Value) UnmarshalJSON(b []byte) error {
	if b == nil || len(b) == 0 {
		return errors.New("Cannot unmarshal nil or empty value")
	}
	rdr := bytes.NewReader(b)
	var err error
	var kind rune
	var content strings.Builder
	var current rune
	for i := 0; err == nil; i++ {
		var char rune
		char, _, err = rdr.ReadRune()
		if i == 0 && char != rune('"') {
			return errors.New("Expected a string")
		}
		if i == 1 {
			kind = char
		}
		if i == 2 && char != rune(':') {
			return errors.New("Expected a string")
		}
		if i > 2 {
			_, err := content.WriteRune(char)
			if err != nil {
				return err
			}
		}
		current = char
	}
	if current != rune('"') {
		return errors.New("Unterminated string")
	}
	_ = kind
	return errors.New("Unable to unmarshal value " + string(b))
}

// MarshalJSON encode the value in format compatible with haystack's JSON:
// https://www.project-haystack.org/doc/Json
func (v *Value) MarshalJSON() ([]byte, error) {
	var output string
	switch v.kind {
	case haystackTypeBool:
		output = fmt.Sprintf("%v", v.b)
	case haystackTypeGrid:
		return json.Marshal(v.g)
	case haystackTypeStr:
		output = `"s:` + *v.str + `"`
	case haystackTypeRef:
		output = `"r:` + string(*v.ref) + `"`
	case haystackTypeRemove:
		output = `"-:"`
	case haystackTypeMarker:
		output = `"m:"`
	case haystackTypeNA:
		output = `"z:"`
	case haystackTypeCoord:
		output = fmt.Sprintf(`"c:%v,%v"`, v.coord.lat, v.coord.long)
	case haystackTypeDate:
		output = `"d:` + v.t.Format("2006-01-02") + `"`
	case haystackTypeTime:
		output = `"h:` + v.t.Format("15:04:05") + `"`
	case haystackTypeDateTime:
		output = `"t:` + v.t.Format(time.RFC3339) + `"`
	case haystackTypeURI:
		output = `"u:` + (*v.u).String() + `"`
	case haystackTypeNumber:
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
		kind: haystackTypeRef,
		ref:  r,
	}
}

// NewStr new string value
func NewStr(s string) *Value {
	return &Value{
		kind: haystackTypeStr,
		str:  &s,
	}
}

// MarkerValue ...
var MarkerValue = &Value{
	kind: haystackTypeMarker,
}

// GetString value; returns an error if the underlying type is not an haystack string
func (v *Value) GetString() (string, error) {
	if v.kind != haystackTypeStr {
		return "", errors.New("value type is not a string")
	}
	return *v.str, nil
}

const (
	// haystackTypeUndefined ...
	haystackTypeUndefined kind = iota
	// haystackTypeGrid is a Grid object
	haystackTypeGrid
	// haystackTypeList Array
	haystackTypeList
	// haystackTypeDict Object
	haystackTypeDict
	// haystackTypenull null
	haystackTypenull
	// haystackTypeBool Boolean
	haystackTypeBool
	// haystackTypeMarker "m:"
	haystackTypeMarker
	// haystackTypeRemove "-:"
	haystackTypeRemove
	// haystackTypeNA "z:"
	haystackTypeNA
	// haystackTypeNumber "n:<float> [unit]" "n:45.5" "n:73.2 °F" "n:-INF"
	haystackTypeNumber
	// haystackTypeRef "r:<id> [dis]"  "r:abc-123" "r:abc-123 RTU #3"
	haystackTypeRef
	// haystackTypeStr "hello" "s:hello"
	haystackTypeStr
	// haystackTypeDate "d:2014-01-03"
	haystackTypeDate
	// haystackTypeTime "h:23:59:00"
	haystackTypeTime
	// haystackTypeDateTime "t:2015-06-08T15:47:41-04:00 New_York"
	haystackTypeDateTime
	// haystackTypeURI "u:http://project-haystack.org/"
	haystackTypeURI
	// haystackTypeCoord "c:<lat>,<lng>" "c:37.545,-77.449"
	haystackTypeCoord
	//haystackTypeXStr "x:Type:value"
	haystackTypeXStr
	// haystackLastType ...
	haystackLastType
)
