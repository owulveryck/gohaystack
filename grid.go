package gohaystack

import (
	"errors"
)

// Grid is a simple database structure, column based
type Grid struct {
	Meta     map[string]string
	labels   []*Label // All the tags used in the grid accrodd various entities
	entities []*Entity
}

// NewGrid creates an empty grid
func NewGrid() *Grid {
	return &Grid{
		Meta: map[string]string{
			"Ver": "3.0",
		},
		labels:   make([]*Label, 0),
		entities: make([]*Entity, 0),
	}
}

// GetEntities from the grid
func (g *Grid) GetEntities() []*Entity {
	return g.entities
}

// Label of a tag. Label is a string and for efficiency we are using pointer to Label in the entities
type Label struct {
	Value   string
	Display string
}

// HaystackID is a Unique identifier for an entity.
// see https://www.project-haystack.org/doc/TagModel#Id
type HaystackID string

// NewHaystackID from string s
func NewHaystackID(s string) *HaystackID {
	id := HaystackID(s)
	return &id
}

// NewLabel from string s
func NewLabel(s string) *Label {
	return &Label{
		Value: s,
	}
}

// NewEntity identified by id. it returns an error if the ID is already taken.
// It is the responsibility of the caller to ensure that the ID does not exists yet in the grid.
func (g *Grid) NewEntity(id *HaystackID) *Entity {
	entity := &Entity{
		id:   id,
		tags: make(map[*Label]*Value),
	}
	g.entities = append(g.entities, entity)
	return entity
}

// AddLabel to the grid; returns nil
// TODO: returns an error if the label already exits
func (g *Grid) AddLabel(l *Label) error {
	if l == nil {
		return errors.New("Cannot add nil label")
	}
	g.labels = append(g.labels, l)
	return nil
}

// GetLabels used in the grid
func (g *Grid) GetLabels() []*Label {
	return g.labels
}
