package gohaystack

// Grid is a simple database structure, column based
type Grid struct {
	Meta     map[string]string
	entities []*Entity
}

// NewGrid creates an empty grid
func NewGrid() *Grid {
	return &Grid{
		Meta: map[string]string{
			"ver": "3.0",
		},
		entities: make([]*Entity, 0),
	}
}

// GetEntityByID returns the entity identified by id, or nil if nothing is found
// if a grid has two elements with the same ID (there is no fence to protect from that),
// only the first one is returned
func (g *Grid) GetEntityByID(id *HaystackID) *Entity {
	for _, entity := range g.entities {
		if entity.GetID() == id {
			return entity
		}
	}
	return nil
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
