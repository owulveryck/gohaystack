package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/owulveryck/gohaystack"
	"gonum.org/v1/gonum/graph/simple"
)

func sampleGrid() *gohaystack.Grid {
	buf := bytes.NewBufferString(carytown)
	g := gohaystack.NewGrid()
	dec := json.NewDecoder(buf)
	err := dec.Decode(&g)
	if err != nil {
		panic(err)
	}
	return g
}

func Test_graphHandler_addNodes(t *testing.T) {
	type fields struct {
		graph *simple.UndirectedGraph
		grid  *gohaystack.Grid
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"simple",
			fields{
				graph: simple.NewUndirectedGraph(),
				grid:  sampleGrid(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gh := &graphHandler{
				graph: tt.fields.graph,
				grid:  tt.fields.grid,
			}
			if err := gh.addNodes(); (err != nil) != tt.wantErr {
				t.Errorf("graphHandler.addNodes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
