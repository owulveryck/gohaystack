package gohaystack

import (
	"errors"
	"net/url"
	"time"
)

type kind int

// Value is an haystack value
type Value struct {
	kind   kind
	str    *string
	number struct {
		value float32
		unit  string
	}
	t     *time.Time
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

	return errors.New("Unable to unmarshal value " + string(b))
}

// MarshalJSON encode the value in format compatible with haystack's JSON:
// https://www.project-haystack.org/doc/Json
func (v *Value) MarshalJSON() ([]byte, error) {
	var output string
	switch v.kind {
	case haystackTypeStr:
		output = `"s:` + *v.str + `"`
	case haystackTypeRef:
		output = `"r:` + string(*v.ref) + `"`
	case haystackTypeMarker:
		output = `"m:"`
	case haystackTypeURI:
		output = `"u:` + (*v.u).String() + `"`
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
