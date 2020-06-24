package gohaystack

import (
	"reflect"
	"testing"
)

func TestGrid_MarshalZinc(t *testing.T) {
	//	blabla := "blabla"
	myID := NewHaystackID("myid")
	type fields struct {
		Meta     map[string]string
		entities []*Entity
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"empty grid",
			fields{},
			nil,
			true,
		},
		{
			"bad meta",
			fields{
				Meta: map[string]string{
					"version": "3.0",
				},
			},
			nil,
			true,
		},
		{
			"good meta, bad version",
			fields{
				Meta: map[string]string{
					"ver": "4.0",
				},
			},
			nil,
			true,
		},
		{
			"good meta, bar Version",
			fields{
				Meta: map[string]string{
					"Ver": "4.0",
				},
			},
			nil,
			true,
		},
		{
			"good meta, good Version",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
			},
			[]byte(`ver:"3.0"\n`),
			false,
		},
		{
			"good meta, good Version, more meta",
			fields{
				Meta: map[string]string{
					"ver":      "3.0",
					"database": "test",
				},
			},
			[]byte(`ver:"3.0" database:"test"\n`),
			false,
		},
		{
			"one column",
			fields{
				Meta: map[string]string{
					"ver":      "3.0",
					"database": "test",
				},
				entities: []*Entity{
					{
						id: myID,
						tags: map[*Label]*Value{
							{Value: "blabla"}: {
								kind: HaystackLastType,
							},
						},
					},
				},
			},
			[]byte(`ver:"3.0" database:"test"\nblabla\n\n`),
			false,
		},
		{
			"one column and display",
			fields{
				Meta: map[string]string{
					"ver":      "3.0",
					"database": "test",
				},
				entities: []*Entity{
					{
						id: myID,
						tags: map[*Label]*Value{
							{
								Value:   "blabla",
								Display: "display",
							}: {
								kind: HaystackLastType,
							},
						},
					},
				},
			},
			[]byte(`ver:"3.0" database:"test"\nblabla dis:"display"\n\n`),
			false,
		},
		{
			"two columns and display",
			fields{
				Meta: map[string]string{
					"ver":      "3.0",
					"database": "test",
				},
				entities: []*Entity{
					{
						id: myID,
						tags: map[*Label]*Value{
							{
								Value: "a",
							}: {
								kind: HaystackLastType,
							},
							{
								Value: "a",
							}: {
								kind: HaystackLastType,
							},
						},
					},
				},
			},
			[]byte(`ver:"3.0" database:"test"\na,a\n,\n`), // This should not be valid, but in the grid the two labels are differents
			false,
		},
		{
			"bad valud",
			fields{
				Meta: map[string]string{
					"ver":      "3.0",
					"database": "test",
				},
				entities: []*Entity{
					{
						id: myID,
						tags: map[*Label]*Value{
							{
								Value: "a",
							}: nil,
							{
								Value: "a",
							}: {
								kind: HaystackLastType,
							},
						},
					},
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			got, err := g.MarshalZinc()
			if (err != nil) != tt.wantErr {
				t.Errorf("Grid.MarshalZinc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.MarshalZinc() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
