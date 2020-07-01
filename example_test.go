package gohaystack

import (
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
}
func ExampleMarshalZinc() {
	g := NewGrid()
	myTagLabel := NewLabel("mytag")
	myTagSite := NewLabel("site")
	myTagLabel.Display = "the display"
	mySite := NewHaystackID("myreference")
	entity := g.NewEntity(mySite)
	myTagValue := NewStr("foo")
	entity.SetTag(myTagLabel, myTagValue)
	entity.SetTag(myTagSite, MarkerValue)
	err := g.MarshalZinc(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
