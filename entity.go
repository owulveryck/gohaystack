package gohaystack


// An Entity is an abstraction for some physical object in the real world.
// Entity carries its tags and an ID
type Entity struct {
	id   *HaystackID // id is private because it should be immutable
	Dis  string      // Dis is public because it is mutable
	tags map[*Label]*Value
}

// SetTag for the entity. No check is done and any existing label is silently overwritten.
func (e *Entity) SetTag(l *Label, v *Value) {
	e.tags[l] = v
}

// GetID of the entity
func (e *Entity) GetID() *HaystackID {
	return e.id
}