package gohaystack

import (
	"encoding/json"
	"errors"
)

// UnmarshalJSON turns a JSON encoded grid into a Grid object
func (g *Grid) UnmarshalJSON(b []byte) error {
	var temp haystackJSONStructure
	return json.Unmarshal(b, &temp)
}

// MarshalJSON in haystack compatible format
func (g *Grid) MarshalJSON() ([]byte, error) {
	var hasVer bool
	var version string
	if v, ok := g.Meta["Ver"]; ok {
		hasVer = true
		version = v
	}
	if v, ok := g.Meta["ver"]; ok {
		hasVer = true
		version = v
	}
	if !hasVer {
		return nil, errors.New("Bad formatting, missing version tag")
	}
	if version != "3.0" {
		return nil, errors.New("Unsupported version " + version)
	}
	carrier := haystackJSONStructure{}
	carrier.Meta = g.Meta
	labels := make(map[*Label]struct{}, 0)
	for _, entity := range g.entities {
		for label := range entity.tags {
			labels[label] = struct{}{}
		}
	}
	carrier.Cols = make([]haystackJSONCol, len(labels))

	carrier.Rows = make([]map[string]*Value, len(g.entities))
	i := 0
	for label := range labels {
		carrier.Cols[i] = haystackJSONCol{
			Name: label.Value,
			Dis:  label.Display,
		}
		i++
	}
	for i, entity := range g.entities {
		carrier.Rows[i] = make(map[string]*Value, len(entity.tags)+1) // all tags + id
		carrier.Rows[i]["id"] = &Value{
			kind: haystackTypeRef,
			ref:  entity.id,
		}
		for k, v := range entity.tags {
			carrier.Rows[i][k.Value] = v
		}
	}

	return json.Marshal(carrier)
}

type haystackJSONCol struct {
	Name string `json:"name"`
	Dis  string `json:"dis,omitempty"`
}
type haystackJSONStructure struct {
	Meta map[string]string   `json:"meta"`
	Cols []haystackJSONCol   `json:"cols"`
	Rows []map[string]*Value `json:"rows"`
}
