package gohaystack

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
)

func TestGrid_UnmarshalJSON(t *testing.T) {
	blabla := "blabla"
	myID := NewHaystackID("myid")
	myID2 := NewHaystackID("myid2")
	type fields struct {
		Meta     map[string]string
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
			"bad json structure",
			fields{},
			args{
				[]byte(`blabla`),
			},
			true,
		},
		{
			"No version",
			fields{},
			args{
				[]byte(`{"meta":{"verion":"3.0"}, "cols":[{"name":"blabla"}],"rows":[{"id":"r:myid","blabla":"s:blabla"}]}`),
			},
			true,
		},
		{
			"bad ver number",
			fields{},
			args{
				[]byte(`{"meta":{"ver":"WRONG"}, "cols":[{"name":"blabla"}],"rows":[{"id":"r:myid","blabla":"s:blabla"}]}`),
			},
			true,
		},
		{
			"bad Ver number",
			fields{},
			args{
				[]byte(`{"meta":{"Ver":"WRONG"}, "cols":[{"name":"blabla"}],"rows":[{"id":"r:myid","blabla":"s:blabla"}]}`),
			},
			true,
		},
		{
			"missing id",
			fields{},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name":"blabla"}],"rows":[{"idd":"r:myid","blabla":"s:blabla"}]}`),
			},
			true,
		},
		{
			"id is not ref",
			fields{},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name":"blabla"}],"rows":[{"id":"s:myid","blabla":"s:blabla"}]}`),
			},
			true,
		},
		{
			"id ok",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
				entities: []*Entity{
					{
						id: NewHaystackID("myid"),
						/*
							tags: map[*Label]*Value{
								&Label{Value: "blabla"}: &Value{
									kind: haystackTypeStr,
									str:  &blabla,
								},
							},
						*/
					},
				},
			},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[],"rows":[{"id":"r:myid"}]}`),
			},
			false,
		},
		{
			"id ok and one entity",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
				entities: []*Entity{
					{
						id: NewHaystackID("myid"),
						tags: map[*Label]*Value{
							{Value: "blabla"}: {
								kind: HaystackTypeStr,
								str:  &blabla,
							},
						},
					},
				},
			},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name": "blabla"}],"rows":[{"id":"r:myid","blabla":"blabla"}]}`),
			},
			false,
		},
		{
			"id ok and two entities with references",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
				entities: []*Entity{
					{
						id: myID,
						tags: map[*Label]*Value{
							{Value: "blabla"}: {
								kind: HaystackTypeStr,
								str:  &blabla,
							},
						},
					},
					{
						id: myID2,
						tags: map[*Label]*Value{
							{Value: "somethingRef"}: {
								kind: HaystackTypeRef,
								ref:  myID,
							},
						},
					},
				},
			},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name": "blabla"},{"name": "somethingRef"}],"rows":[{"id":"r:myid","blabla":"blabla"},{"id":"r:myid2","somethingRef":"r:myid"}]}`),
			},
			false,
		},

		{
			"Undeclared label",
			fields{},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name": "blablabla"}],"rows":[{"id":"r:myid","blabla":"blabla"}]}`),
			},
			true,
		},
		/*
			{
				"simple",
				fields{
					Meta: map[string]string{
						"ver": "3.0",
					},
					entities: []*Entity{
						&Entity{
							id: NewHaystackID("myid"),
							tags: map[*Label]*Value{
								&Label{Value: "blabla"}: &Value{
									kind: haystackTypeStr,
									str:  &blabla,
								},
							},
						},
					},
				},
				args{
					[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name":"blabla"}],"rows":[{"id":"r:myid","blabla":"s:blabla"}]}`),
				},
				false,
			},
				{
					"sample",
					fields{
						Meta: map[string]string{
							"ver": "3.0",
						},
						//entities: make([]*Entity, 0),
					},
					args{
						samplePayload,
					},
					false,
				},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			gg := new(Grid)
			if err := gg.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Grid.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(g.Meta, gg.Meta) {
				t.Errorf("Grid.UnmarshalJSON() = %v, want %v", gg.Meta, g.Meta)
			}
			if len(g.entities) != len(gg.entities) {
				t.Errorf("Grid.UnmarshalJSON() = found %v entities, want %v", len(gg.entities), len(g.entities))
			}
			for i := 0; i < len(g.entities); i++ {
				gentity := g.entities[i]
				ggentity := gg.entities[i]
				if *gentity.id != *ggentity.id {
					t.Errorf("Grid.UnmarshalJSON() = bad id %v, want %v", gentity.id, ggentity.id)

				}
				if len(gentity.tags) != len(ggentity.tags) {
					t.Errorf("Grid.UnmarshalJSON() = found %v tags, want %v", len(ggentity.tags), len(gentity.tags))
				}
				for k, v := range ggentity.tags {
					var vv *Value
					var ok bool
					if vv, ok = ggentity.tags[k]; !ok {
						t.Errorf("Grid.UnmarshalJSON() = expected tag %v not found", k)
					}
					if !reflect.DeepEqual(v, vv) {
						t.Errorf("Grid.UnmarshalJSON() = bad vallue for tag %v found %v tags, want %v", k, v, vv)
					}
				}
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
				entities: []*Entity{
					{
						id: myid,
						tags: map[*Label]*Value{
							siteLabel:   MarkerValue,
							blablaLabel: blablaValue,
						},
					},
				},
			},
			[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name":"site"},{"name":"blabla"}],"rows":[{"id":"r:myid","blabla":"s:blabla","site":"m:"}]}`),
			false,
		},
		{
			"simple example with dis",
			fields{
				Meta: map[string]string{
					"ver": "3.0",
				},
				entities: []*Entity{
					{
						id: myid,
						tags: map[*Label]*Value{
							blablaLabelWithDis: blablaValue,
						},
					},
				},
			},
			[]byte(`{
				"meta":{
				   "ver":"3.0"
				},
				"cols":[
				   {
					  "name":"blabla",
					  "dis": "blabla display"
				   }
				],
				"rows":[
				   {
					  "blabla":"s:blabla",
					  "id":"r:myid"
				   }
				]
			 }`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			got, err := g.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Grid.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {

				var a, b haystackJSONStructureTest
				err = json.Unmarshal(tt.want, &a)
				if err != nil {
					t.Fatal(err)
				}
				err = json.Unmarshal(got, &b)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(a.Meta, b.Meta) {
					t.Errorf("Grid.MarshalJSON() = %v, want %v", b, a)
				}
				if !reflect.DeepEqual(a.Rows, b.Rows) {
					t.Errorf("Grid.MarshalJSON() = %v, want %v", b, a)
				}
				sort.Sort(labelsByAlphabeticalOrder(a.Cols))
				sort.Sort(labelsByAlphabeticalOrder(b.Cols))
				if !reflect.DeepEqual(a.Cols, b.Cols) {
					t.Errorf("Grid.MarshalJSON() = %v, want %v", b.Cols, a.Cols)
				}

			}
		})
	}
}

type labelsByAlphabeticalOrder []haystackJSONColTest

func (a labelsByAlphabeticalOrder) Len() int           { return len(a) }
func (a labelsByAlphabeticalOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a labelsByAlphabeticalOrder) Less(i, j int) bool { return a[i].Name < a[j].Name }

type haystackJSONColTest struct {
	Name string `json:"name"`
	Dis  string `json:"dis,omitempty"`
}
type haystackJSONStructureTest struct {
	Meta map[string]string     `json:"meta"`
	Cols []haystackJSONColTest `json:"cols"`
	Rows []map[string]string   `json:"rows"`
}
