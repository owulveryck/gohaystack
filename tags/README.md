# package tags

package tags holds pre-instantiated  variables with official tags of the haystack-project

## Example

```go
func ExampleMarker() {
	g := gohaystack.NewGrid()
	id := gohaystack.NewHaystackID("myid")
	firstEntity := g.NewEntity(id)
	secondEntity := g.NewEntity(id)
	//  Set the Site tag as a marker tag
    firstEntity.SetTag(Site())
    // or do it with the label
    firstEntity.SetTag(SiteLabel, gohaystack.MarkerValue)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(g)
}
```