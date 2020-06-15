package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"encoding/json"

	"github.com/owulveryck/gohaystack"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

var maxnodes = flag.Int("maxnodes", 0, "maximum nodes to add to the graph (0 for all)")
var ports = flag.Bool("ports", false, "generate Graphviz with ports links")

func main() {
	flag.Parse()
	g := gohaystack.NewGrid()
	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(&g)
	if err != nil {
		log.Fatal(err)
	}
	gr := newGraphHandler(g)
	err = gr.addNodes()
	if err != nil {
		log.Fatal(err)
	}
	gr.addEdges()

	result, err := dot.Marshal(gr.graph, "", "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(result))

}

type node struct {
	id      int64
	content *gohaystack.Entity
}

// node's uniq ID to fulfil graph.Node
func (n *node) ID() int64 {
	return n.id
}

// Attributes is for graphviz output. It specifies the "label" of the node (a table)
func (n *node) Attributes() []encoding.Attribute {
	attrs := []encoding.Attribute{
		{
			Key:   "id",
			Value: fmt.Sprintf(`"%v"`, n.id),
		},
		{
			Key:   "shape",
			Value: "Mrecord",
		},
		{
			Key:   "label",
			Value: fmt.Sprintf(`%v`, n.generateLabel()),
		},
	}
	return attrs
}

type graphHandler struct {
	graph *simple.UndirectedGraph
	grid  *gohaystack.Grid
	refs  []*gohaystack.Label // Hold all of the labels sometingRef
}

func newGraphHandler(g *gohaystack.Grid) *graphHandler {
	return &graphHandler{
		graph: simple.NewUndirectedGraph(),
		grid:  g,
		refs:  make([]*gohaystack.Label, 0),
	}
}

func (n *node) generateLabel() string {
	output := "{"
	output = fmt.Sprintf("%v{%v}|", output, *n.content.GetID())
	var tmp string
	for k, v := range n.content.GetTags() {
		if v != nil {
			display, _ := json.Marshal(v)
			if tmp != "" {
				tmp = tmp + "|"
			}
			switch v.GetKind() {
			/*
					case gohaystack.HaystackTypeRef:
						tmp = tmp + fmt.Sprintf("{<%v>%v|@%v}", cols[i], cols[i], v.Value)
					case gohaystack.HaystackTypeMarker:
						tmp = tmp + fmt.Sprintf("{%v}", cols[i])
				case gohaystack.HaystackTypeRef:
					tmp = tmp + fmt.Sprintf("{<%v>%v|%@v}", k.Value, *k, string(display))
			*/
			default:
				tmp = tmp + fmt.Sprintf("{%v|%v}", *k, string(display))
			}
		}
	}
	output = output + tmp + "}"
	return output
}

type edgeWithPorts struct {
	simple.Edge
	fromPort, toPort string
}

func (e edgeWithPorts) ReversedEdge() graph.Edge {
	e.F, e.T = e.T, e.F
	e.fromPort, e.toPort = e.toPort, e.fromPort
	return e
}

func (e edgeWithPorts) FromPort() (string, string) {
	return e.fromPort, ""
}

func (e edgeWithPorts) ToPort() (string, string) {
	return e.toPort, ""
}

func newEdge(from, to graph.Node, fromPort, toPort string) edgeWithPorts {
	return edgeWithPorts{
		Edge: simple.Edge{
			F: from,
			T: to,
		},
		fromPort: fromPort,
		toPort:   toPort,
	}
}

func (gh *graphHandler) addNodes() error {
	entities := gh.grid.GetEntities()
	ids := make(map[*gohaystack.HaystackID]int64, len(entities))
	for i, entity := range entities {
		n := &node{
			id:      int64(i),
			content: entity,
		}
		ids[entity.GetID()] = int64(i)
		gh.graph.AddNode(n)
	}
	// Now generate the edges
	/*
		for i, entity := range entities {
		}
	*/
	return nil
}

func (gh *graphHandler) addEdges() error {
	return nil
}
