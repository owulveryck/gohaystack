package gohaystack

import (
	"encoding/json"
	"log"
	"os"
)

func ExampleGrid() {
	g := NewGrid()
	myTagLabel := NewLabel("mytag")
	myTagSite := NewLabel("site")
	myTagLabel.Display = "the display"
	mySite := NewHaystackID("myreference")
	entity := g.NewEntity(mySite)
	myTagValue := NewStr("foo")
	entity.SetTag(myTagLabel, myTagValue)
	entity.SetTag(myTagSite, MarkerValue)
	// Do something with the grid ... :D
	// ...
	// Encode it in json
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(g)
	if err != nil {
		log.Fatal(err)
	}

}
