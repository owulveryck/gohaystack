package tags

import (
	"encoding/json"
	"os"

	"github.com/owulveryck/gohaystack"
)

func ExampleMarker() {
	g := gohaystack.NewGrid()
	id := gohaystack.NewHaystackID("myid")
	entity := g.NewEntity(id)
	//  Set the Site tag as a marker tag
	entity.SetTag(Site())
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(g)
	// output: {"meta":{"Ver":"3.0"},"cols":[{"name":"site"}],"rows":[{"id":"r:myid","site":"m:"}]}
}
