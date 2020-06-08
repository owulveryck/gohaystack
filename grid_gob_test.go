package gohaystack

import (
	"bytes"
	"encoding/gob"
	"net/url"
	"reflect"
	"testing"
)

func TestGrid_GobEncode(t *testing.T) {
	g := NewGrid()
	g.AddColumn("col1", "col1")
	g.AddColumn("col2", "col2")
	g.NewRow()
	g.Set(0, "col1", &TypedValue{
		Value: "bla",
	})
	g.Set(0, "col2", &TypedValue{
		Value: 42,
	})
	g.NewRow()
	g.Set(1, "col2", &TypedValue{
		Value: 42.42,
	})
	g.NewRow()
	myURL, _ := url.Parse("https://example.com")
	g.Set(1, "col2", &TypedValue{
		Type:  HaystackTypeURI,
		Value: myURL,
	})
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(g)
	if err != nil {
		t.Fatal(err)
	}
	var result *Grid
	dec := gob.NewDecoder(&b)
	err = dec.Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	if result.lastCol != g.lastCol {
		t.Fail()
	}
	if result.numberOfRows != g.numberOfRows {
		t.Fail()
	}
	if !reflect.DeepEqual(result.colsDis, g.colsDis) {
		t.Fail()
	}
	if !reflect.DeepEqual(result.Meta, g.Meta) {
		t.Fail()
	}
	if !reflect.DeepEqual(result.Cols, g.Cols) {
		t.Fail()
	}
	if len(result.db) != len(g.db) {
		t.Fail()
	}
	for k, v := range g.db {
		var ok bool
		var val interface{}
		if val, ok = result.db[k]; !ok {
			t.Log("key not found")
			t.Fail()
		}
		if !reflect.DeepEqual(v, val) {
			t.Logf("%v != $%v", v, val)
			t.Fail()
		}
	}
}
