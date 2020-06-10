package gohaystack

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGrid_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Meta     map[string]string
		labels   []*Label
		entities []*Entity
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"sample",
			fields{
				Meta: map[string]string{
					"Ver": "3,0",
				},
				labels:   make([]*Label, 0),
				entities: make([]*Entity, 0),
			},
			args{
				samplePayload,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				labels:   tt.fields.labels,
				entities: tt.fields.entities,
			}
			if err := g.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Grid.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGrid_MarshalJSON(t *testing.T) {
	siteLabel := NewLabel("site")
	blablaLabel := NewLabel("blabla")
	blablaLabelWithDis := NewLabel("blabla")
	blablaLabelWithDis.Display = "blabla display"
	blablaValue := NewStr("blabla")
	myid := NewHaystackID("myid")
	type fields struct {
		Meta     map[string]string
		labels   []*Label
		entities []*Entity
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"Missing version",
			fields{
				Meta: map[string]string{
					"Version": "3.0",
				},
				labels:   []*Label{siteLabel, blablaLabel},
				entities: []*Entity{},
			},
			nil,
			true,
		},
		{
			"Bad version Ver",
			fields{
				Meta: map[string]string{
					"Ver": "4.0",
				},
				labels:   []*Label{siteLabel, blablaLabel},
				entities: []*Entity{},
			},
			nil,
			true,
		},
		{
			"Bad version ver",
			fields{
				Meta: map[string]string{
					"ver": "4.0",
				},
				labels:   []*Label{siteLabel, blablaLabel},
				entities: []*Entity{},
			},
			nil,
			true,
		},
		{
			"simple example",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
				labels: []*Label{siteLabel, blablaLabel},
				entities: []*Entity{
					{
						id: myid,
						tags: map[*Label]*Value{
							blablaLabel: blablaValue,
						},
					},
				},
			},
			[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name":"site"},{"name":"blabla"}],"rows":[{"id":"r:myid","blabla":"s:blabla"}]}`),
			false,
		},
		{
			"simple example with dis",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
				labels: []*Label{siteLabel, blablaLabelWithDis},
				entities: []*Entity{
					{
						id: myid,
						tags: map[*Label]*Value{
							blablaLabel: blablaValue,
						},
					},
				},
			},
			[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name":"site"},{"name":"blabla","dis":"blabla display"}],"rows":[{"id":"r:myid","blabla":"s:blabla"}]}`),
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
			got, err := g.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Grid.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var a, b map[string]interface{}
				err = json.Unmarshal(tt.want, &a)
				if err != nil {
					t.Fatal(err)
				}
				err = json.Unmarshal(got, &b)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(a, b) {
					t.Errorf("Grid.MarshalJSON() = %v, want %v", a, b)
				}
			}
		})
	}
}
