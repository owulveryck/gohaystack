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
				labels:   make([]*Label, 0),
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
				labels:   tt.fields.labels,
				entities: tt.fields.entities,
			}
			if got := g.NewEntity(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.NewEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_AddLabel(t *testing.T) {
	type fields struct {
		Meta     map[string]string
		labels   []*Label
		entities []*Entity
	}
	type args struct {
		l *Label
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Add nil label",
			fields{
				Meta: map[string]string{
					"Ver": "3,0",
				},
				labels:   make([]*Label, 0),
				entities: make([]*Entity, 0),
			},
			args{
				nil,
			},
			true,
		},
		{
			"Add non-nil empty label",
			fields{
				Meta: map[string]string{
					"Ver": "3,0",
				},
				labels:   make([]*Label, 0),
				entities: make([]*Entity, 0),
			},
			args{
				NewLabel(""),
			},
			false,
		},
		{
			"Add non-nil non-empty label",
			fields{
				Meta: map[string]string{
					"Ver": "3,0",
				},
				labels:   make([]*Label, 0),
				entities: make([]*Entity, 0),
			},
			args{
				NewLabel("label"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				labels:   tt.fields.labels,
				entities: tt.fields.entities,
			}
			if err := g.AddLabel(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Grid.AddLabel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && g.labels[len(g.labels)-1] != tt.args.l {
				t.Errorf("Grid.AddLabel() error label not added")
			}
		})
	}
}

func TestGrid_GetLabels(t *testing.T) {
	label1 := NewLabel("label1")
	label2 := NewLabel("label2")

	type fields struct {
		Meta     map[string]string
		labels   []*Label
		entities []*Entity
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Label
	}{
		{
			"simple test",
			fields{
				Meta: map[string]string{
					"Ver": "3,0",
				},
				labels:   []*Label{label1, label2, nil},
				entities: make([]*Entity, 0),
			},
			[]*Label{label1, label2, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				labels:   tt.fields.labels,
				entities: tt.fields.entities,
			}
			if got := g.GetLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.GetLabels() = %v, want %v", got, tt.want)
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
				labels:   tt.fields.labels,
				entities: tt.fields.entities,
			}
			if got := g.GetEntities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}
