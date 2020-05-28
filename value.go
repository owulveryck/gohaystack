package gohaystack

import (
	"fmt"
	"reflect"
	"strings"
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

// Hash ...
func (tv *TypedValue) Hash() string {
	return fmt.Sprintf("%v/%v", tv.Type, tv.Value)
}

// TODO ...
func inferType(value interface{}) (HaystackType, interface{}) {
	if _, ok := value.(bool); ok {
		return HaystackTypeBool, value
	}
	if valueString, ok := value.(string); ok {
		vals := strings.Split(valueString, ":")
		if len(vals) > 1 {
			if len(vals[0]) != 1 {
				return HaystackTypeUndefined, value
			}
			switch vals[0] {
			case "r":
				return HaystackTypeRef, strings.Join(vals[1:], ":")
			case "m":
				return HaystackTypeBool, true
			default:
				return HaystackTypeUndefined, strings.Join(vals[1:], ":")
			}
		}
	}
	return HaystackTypeUndefined, value
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
