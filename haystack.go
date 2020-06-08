package gohaystack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
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
	case HaystackTypeURI:
		return fmt.Sprintf(`"u:%v"`, v.Value.(*url.URL).String())
	case HaystackTypeTime:
		return fmt.Sprintf(`"h:%v"`, v.Value.(time.Time).Format("15:04:05"))
	case HaystackTypeDate:
		return fmt.Sprintf(`"d:%v"`, v.Value.(time.Time).Format("2006-01-02"))
	case HaystackTypeDateTime:
		return fmt.Sprintf(`"t:%v"`, v.Value.(time.Time).Format(time.RFC3339))
	case HaystackTypeRef:
		return fmt.Sprintf(`"r:%v"`, v.Value)
	case HaystackTypeNumber:
		var unit string
		if v.Value.(*HaystackNumber).Unit != "" {
			unit = " " + v.Value.(*HaystackNumber).Unit
		}
		return fmt.Sprintf(`"n:%v%v"`, v.Value.(*HaystackNumber).Value, unit)
	case HaystackTypeMarker:
		return fmt.Sprintf(`"m:"`)
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
