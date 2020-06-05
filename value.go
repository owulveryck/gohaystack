package gohaystack

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// NewTypedValue holding the haystack type (not tag).
// There is not check done to see if the interface{} underlying type is compatible
// with the haystack type.
func NewTypedValue(typ HaystackType, value interface{}) *TypedValue {
	if typ >= HaystackLastType {
		typ = HaystackTypeUndefined
	}
	return &TypedValue{
		Type:  typ,
		Value: value,
	}
}

// TypedValue is any value that can be associated with a HaystackType
type TypedValue struct {
	Type  HaystackType
	Value interface{}
}

// HaystackNumber is an integer or floating point number annotated with an optional unit of measurement.
type HaystackNumber struct {
	Value float32
	Unit  string
}

// Hash ...
func (tv *TypedValue) Hash() string {
	return fmt.Sprintf("%v/%v", tv.Type, tv.Value)
}

// TODO ...
func inferType(value interface{}) (HaystackType, interface{}, error) {
	if _, ok := value.(bool); ok {
		return HaystackTypeBool, value, nil
	}
	valueString, ok := value.(string)

	if !ok {
		return HaystackTypeStr, value, nil
	}
	vals := strings.Split(valueString, ":")
	if len(vals) > 1 {
		if len(vals[0]) != 1 {
			return HaystackTypeUndefined, value, nil
		}
		switch vals[0] {
		case "s":
			return HaystackTypeStr, strings.Join(vals[1:], ":"), nil
		case "n":
			elements := strings.Fields(strings.Join(vals[1:], ":"))
			f, err := strconv.ParseFloat(elements[0], 32)
			if err != nil {
				return HaystackTypeUndefined, nil, err
			}
			var unit string
			if len(elements) > 1 {
				unit = strings.Join(elements[1:], " ")
			}
			return HaystackTypeNumber, &HaystackNumber{
				Value: float32(f),
				Unit:  unit,
			}, nil
		case "r":
			return HaystackTypeRef, strings.Fields(strings.Join(vals[1:], ":"))[0], nil
		case "d":
			t, err := time.Parse("2006-01-02", strings.Join(vals[1:], ":"))
			return HaystackTypeDate, t, err
		case "h":
			t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "0000-01-01 "+strings.Join(vals[1:], ":")+" +0000 UTC")
			return HaystackTypeTime, t, err
		case "t":
			elements := strings.Fields(strings.Join(vals[1:], ":"))
			var err error
			var t time.Time
			/*
				var loc *time.Location
					if len(elements) > 1 {
						loc, err = time.LoadLocation(elements[1])
						if err != nil {
							return HaystackTypeUndefined, nil, err
						}
						t, err = time.ParseInLocation(time.RFC3339, elements[0], loc)
					} else {
			*/
			t, err = time.Parse(time.RFC3339, elements[0])
			/*
				}
			*/
			return HaystackTypeDateTime, t, err
		case "m":
			return HaystackTypeMarker, true, nil
		case "u":
			u, err := url.Parse(strings.Join(vals[1:], ":"))
			return HaystackTypeURI, u, err
		default:
			return HaystackTypeUndefined, strings.Join(vals[1:], ":"), nil
		}
	}
	return HaystackTypeStr, value, nil
}

// Equal returns true if type and value are equal
func (tv *TypedValue) Equal(tv2 *TypedValue) bool {
	if tv == nil && tv2 != tv {
		return false
	}
	if tv.Value == nil && tv.Value == tv2.Value {
		return true
	}
	if !reflect.DeepEqual(tv.Value, tv2.Value) {
		return false
	}
	if !reflect.DeepEqual(tv.Type, tv2.Type) {
		return false
	}
	return true
}
