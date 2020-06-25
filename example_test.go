package gohaystack

import (
	"fmt"
	"log"
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
	b, err := g.MarshalZinc()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
