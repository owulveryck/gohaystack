package gohaystack

import (
	"encoding/json"
	"fmt"
)

type haystack struct {
	Meta map[string]string `json:"meta"`
	Cols []haystackCol     `json:"cols"`
	Rows []haystackRow     `json:"rows"`
}

type haystackCol struct {
	Name string `json:"name"`
	Dis  string `json:"dis,omitempty"`
}

type haystackMeta struct {
	Ver      string `json:"ver"`
	ProjName string `json:"projName"`
}

type haystackRow []haystackKVPair

type haystackKVPair struct {
	Name  string
	Value *TypedValue
}

func (hr haystackRow) MarshalJSON() ([]byte, error) {
	output := `{`

	for _, v := range []haystackKVPair(hr) {
		if v.Value != nil {
			output = fmt.Sprintf("%v%v,", output, v)
		}
	}

	sz := len(output)
	if sz > 0 && output[sz-1] == ',' {
		output = output[:sz-1]
	}
	return []byte(output + `}`), nil
}

func (kv haystackKVPair) String() string {
	return fmt.Sprintf(`"%v": %v`, kv.Name, kv.Value.stringJSON())
}

func (v *TypedValue) stringJSON() string {
	// TODO, switch cases depending of the type
	switch v.Type {
	case HaystackTypeRef:
		return fmt.Sprintf(`"r:%v"`, v.Value)
	case HaystackTypeMarker:
		return fmt.Sprintf(`"m:"`)
	case HaystackTypeNumber:
		return fmt.Sprintf(`"n:%v"`, v.Value)
	case HaystackTypeGrid:
		b, err := json.Marshal(v.Value)
		if err != nil {
			panic(err)
		}
		return string(b)
	default:
		return fmt.Sprintf(`"%v"`, v.Value)
	}
}

// HaystackType represents a type handled by haystack
type HaystackType int

const (
	// HaystackTypeUndefined ...
	HaystackTypeUndefined HaystackType = iota
	// HaystackTypeGrid is a Grid object
	HaystackTypeGrid
	// HaystackTypeList Array
	HaystackTypeList
	// HaystackTypeDict Object
	HaystackTypeDict
	// HaystackTypenull null
	HaystackTypenull
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
	// HaystackTypeTime "h:23:59"
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
