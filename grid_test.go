package gohaystack

import (
	"reflect"
	"testing"
)

func TestNewGrid(t *testing.T) {
	tests := []struct {
		name string
		want *Grid
	}{
		{
			"simple grid",
			&Grid{
				Meta: map[string]string{
					"Ver": "3.0",
				},
				entities: make([]*Entity, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_NewEntity(t *testing.T) {
	nullID := NewHaystackID("")
	type fields struct {
		Meta     map[string]string
		labels   []*Label
		entities []*Entity
	}
	type args struct {
		id *HaystackID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Entity
	}{
		{
			"new Entity with empty ID",
			fields{
				Meta: map[string]string{
					"Ver": "3,0",
				},
				labels:   make([]*Label, 0),
				entities: make([]*Entity, 0),
			},
			args{
				nullID,
			},
			&Entity{
				id:   nullID,
				tags: make(map[*Label]*Value),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			if got := g.NewEntity(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.NewEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_GetEntities(t *testing.T) {
	type fields struct {
		Meta     map[string]string
		labels   []*Label
		entities []*Entity
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Entity
	}{
		{
			"simple test",
			fields{
				entities: []*Entity{},
			},
			[]*Entity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			if got := g.GetEntities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_GetEntityByID(t *testing.T) {
	id1 := NewHaystackID("id1")
	id2 := NewHaystackID("id2")
	id3 := NewHaystackID("id3")
	id4 := NewHaystackID("id1")
	entity1 := &Entity{id: id1}
	entity2 := &Entity{id: id2}
	entity3 := &Entity{id: id3}
	entity4 := &Entity{id: id1}
	type fields struct {
		Meta     map[string]string
		entities []*Entity
	}
	type args struct {
		id *HaystackID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Entity
	}{
		{
			"nil return",
			fields{
				entities: []*Entity{
					entity1,
					entity2,
					entity3,
				},
			},
			args{
				id4,
			},
			nil,
		},
		{
			"one entity match return",
			fields{
				entities: []*Entity{
					entity1,
					entity2,
					entity3,
				},
			},
			args{
				id1,
			},
			entity1,
		},
		{
			"one entity match return",
			fields{
				entities: []*Entity{
					entity1,
					entity2,
					entity3,
					entity4,
				},
			},
			args{
				id1,
			},
			entity1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			if got := g.GetEntityByID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.GetEntityByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
