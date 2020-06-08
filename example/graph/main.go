package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

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
	gridDB := gohaystack.NewGrid()
	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(&gridDB)
	if err != nil {
		log.Fatal(err)
	}
	g := newGraphHandler(gridDB)
	err = g.addNodes()
	if err != nil {
		log.Fatal(err)
	}
	g.addEdges()

	result, err := dot.Marshal(g.graph, "", "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(result))
}

var cols []string

type graphHandler struct {
	nodesCounter int64
	graph        *simple.UndirectedGraph
	grid         *gohaystack.Grid
	colID        int              // the index of the ID within the graph
	refsID       []int            // the indexes of all references *Ref in haystack
	dict         map[string]*node // link between haystack ID and node ID
}

func newGraphHandler(g *gohaystack.Grid) *graphHandler {
	var id int
	var refs []int
	var validRef = regexp.MustCompile(`.*Ref$`)
	cols = make([]string, len(g.Cols))
	for i, name := range g.Cols {
		cols[i] = name
		if name == "id" {
			id = i
		}
	}
	for i, name := range g.Cols {
		if validRef.MatchString(name) {
			refs = append(refs, i)
		}
	}

	return &graphHandler{
		nodesCounter: 0,
		graph:        simple.NewUndirectedGraph(),
		grid:         g,
		colID:        id,
		refsID:       refs,
		dict:         make(map[string]*node),
	}
}

func (gh *graphHandler) addNodes() error {
	it := gohaystack.NewRowIterator(gh.grid)
	for i := 0; it.Next(); i++ {
		if *maxnodes != 0 && i > *maxnodes {
			return nil
		}
		row := it.Row()
		content := make([]*gohaystack.Tag, len(row))
		copy(content, row)
		n := &node{
			id:      gh.nodesCounter,
			content: content,
		}
		gh.graph.AddNode(n)
		gh.dict[row[gh.colID].Hash()] = n
		gh.nodesCounter++
	}
	return nil
}

func (gh *graphHandler) addEdges() error {
	nodesIT := gh.graph.Nodes()
	for nodesIT.Next() {
		currNode := nodesIT.Node()
		row := currNode.(*node).content
		from := gh.dict[row[gh.colID].Hash()]
		for _, i := range gh.refsID {
			if row[i] == nil {
				continue
			}
			to := gh.dict[row[i].Hash()]
			if to != nil && from.ID() != to.ID() {
				if *ports {
					gh.graph.SetEdge(newEdge(from, to, gh.grid.Cols[i], "id"))
				} else {
					gh.graph.SetEdge(newEdge(from, to, "", ""))
				}
			}
		}
	}
	return nil
}

type node struct {
	id      int64
	content []*gohaystack.Tag
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

func (n *node) generateLabel() string {
	//output := fmt.Sprintf("%v|{", n.id)
	output := "{"
	var tmp string
	for i, v := range n.content {
		if v != nil {
			if tmp != "" {
				tmp = tmp + "|"
			}
			switch v.Kind {
			case gohaystack.HaystackTypeRef:
				tmp = tmp + fmt.Sprintf("{<%v>%v|@%v}", cols[i], cols[i], v.Value)
			case gohaystack.HaystackTypeMarker:
				tmp = tmp + fmt.Sprintf("{%v}", cols[i])
			default:
				tmp = tmp + fmt.Sprintf("{%v|%v}", cols[i], v.Value)
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
