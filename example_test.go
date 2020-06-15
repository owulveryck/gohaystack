package gohaystack

import "github.com/kr/pretty"

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
	pretty.Print(g)
	// output: &gohaystack.Grid{
	//     Meta:     {"Ver":"3.0"},
	//     entities: {
	//         &gohaystack.Entity{
	//             id:   &"myreference",
	//             Dis:  "",
	//             tags: {
	//                 &gohaystack.Label{Value:"mytag", Display:"the display"}: &gohaystack.Value{
	//                     kind:   11,
	//                     str:    &"foo",
	//                     number: struct { value float32; unit gohaystack.Unit }{},
	//                     b:      false,
	//                     t:      time.Time{},
	//                     u:      (*url.URL)(nil),
	//                     ref:    (*gohaystack.HaystackID)(nil),
	//                     g:      (*gohaystack.Grid)(nil),
	//                     dict:   {},
	//                     list:   nil,
	//                     coord:  struct { long float32; lat float32 }{},
	//                 },
	//                 &gohaystack.Label{Value:"site", Display:""}: &gohaystack.Value{
	//                     kind:   6,
	//                     str:    (*string)(nil),
	//                     number: struct { value float32; unit gohaystack.Unit }{},
	//                     b:      false,
	//                     t:      time.Time{},
	//                     u:      (*url.URL)(nil),
	//                     ref:    (*gohaystack.HaystackID)(nil),
	//                     g:      (*gohaystack.Grid)(nil),
	//                     dict:   {},
	//                     list:   nil,
	//                     coord:  struct { long float32; lat float32 }{},
	//                 },
	//             },
	//         },
	//     },
	// }
}
