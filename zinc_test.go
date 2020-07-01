package gohaystack

import (
	"bytes"
	"math"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func newString(s string) *string {
	return &s
}

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
			[]byte("ver:\"3.0\"\n"),
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
			[]byte("ver:\"3.0\" database:\"test\"\n"),
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
			[]byte("ver:\"3.0\" database:\"test\"\nid,blabla\n@myid,\n"),
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
			[]byte("ver:\"3.0\" database:\"test\"\nid,blabla dis:\"display\"\n@myid,\n"),
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
			[]byte("ver:\"3.0\" database:\"test\"\nid,a,a\n@myid,,\n"), // This should not be valid, but in the grid the two labels are differents
			false,
		},
		{
			"nil value",
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
			[]byte("ver:\"3.0\" database:\"test\"\nid,a,a\n@myid,"), // This should not be valid, but in the grid the two labels are differents
			true,
		},
		{
			"bad value",
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
								kind: HaystackTypeUndefined,
							},
						},
					},
				},
			},
			[]byte("ver:\"3.0\" database:\"test\"\nid,a,a\n@myid,,"), // This should not be valid, but in the grid the two labels are differents
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:     tt.fields.Meta,
				entities: tt.fields.entities,
			}
			var got bytes.Buffer
			err := g.MarshalZinc(&got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grid.MarshalZinc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Bytes(), []byte(tt.want)) {
				t.Errorf("Grid.MarshalZinc() = %v, want %v", got.Bytes(), tt.want)
			}
		})
	}
}

func TestValue_MarshalZinc(t *testing.T) {
	abc := "abc"
	def := "def"
	date := time.Date(2009, time.January, 02, 15, 04, 05, 06, time.UTC)
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing", secondsEastOfUTC)
	sameTimeInBeijing := time.Date(2009, 1, 2, 23, 04, 05, 06, beijing)

	kwh := NewUnit("kwh")
	uri, _ := url.Parse("https://example.com")
	type fields struct {
		kind   Kind
		str    *string
		number struct {
			value float32
			unit  Unit
		}
		b     bool
		t     time.Time
		u     *url.URL
		ref   *HaystackID
		g     *Grid
		dict  map[string]*Value
		list  []*Value
		coord struct {
			long float32
			lat  float32
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		orwant  []byte
		wantErr bool
	}{
		{
			"nil",
			fields{
				kind: HaystackLastType,
			},
			nil,
			nil,
			false,
		},
		{
			"Undefined",
			fields{
				kind: HaystackTypeUndefined,
			},
			nil,
			nil,
			true,
		},
		{
			"boolean true",
			fields{
				kind: HaystackTypeBool,
				b:    true,
			},
			[]byte(`T`),
			nil,
			false,
		},
		{
			"boolean false",
			fields{
				kind: HaystackTypeBool,
			},
			[]byte(`F`),
			nil,
			false,
		},
		{
			"Remove",
			fields{
				kind: HaystackTypeRemove,
			},
			[]byte(`R`),
			nil,
			false,
		},
		{
			"NA",
			fields{
				kind: HaystackTypeNA,
			},
			[]byte(`NA`),
			nil,
			false,
		},
		{
			"NULL",
			fields{
				kind: HaystackTypeNull,
			},
			[]byte(`N`),
			nil,
			false,
		},
		{
			"Marker",
			fields{
				kind: HaystackTypeMarker,
			},
			[]byte(`M`),
			nil,
			false,
		},
		{
			"string",
			fields{
				kind: HaystackTypeStr,
				str:  newString("bla"),
			},
			[]byte(`"bla"`),
			nil,
			false,
		},
		{
			"xstr",
			fields{
				kind: HaystackTypeXStr,
			},
			nil,
			nil,
			true,
		},
		{
			"number",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 42.0,
				},
			},
			[]byte(`42`),
			nil,
			false,
		},
		{
			"number with unit",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 42.0,
					unit:  kwh,
				},
			},
			[]byte(`42 kwh`),
			nil,
			false,
		},
		{
			"number float",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 42.1,
				},
			},
			[]byte(`42.1`),
			nil,
			false,
		},
		{
			"date",
			fields{
				kind: HaystackTypeDate,
				t:    date,
			},
			[]byte(`2009-01-02`),
			nil,
			false,
		},
		{
			"URI",
			fields{
				kind: HaystackTypeURI,
				u:    uri,
			},
			[]byte("`https://example.com`"),
			nil,
			false,
		},
		{
			"time",
			fields{
				kind: HaystackTypeTime,
				t:    date,
			},
			[]byte(`15:04:05.000000006`),
			nil,
			false,
		},
		{
			"coord",
			fields{
				kind: HaystackTypeCoord,
				coord: struct {
					long float32
					lat  float32
				}{
					42,
					-42,
				},
			},
			[]byte(`C(-42,42)`),
			nil,
			false,
		},
		{
			"datetime",
			fields{
				kind: HaystackTypeDateTime,
				t:    date,
			},
			[]byte(`2009-02-01T15:04:05+0000 UTC`),
			nil,
			false,
		},
		{
			"datetime beijin",
			fields{
				kind: HaystackTypeDateTime,
				t:    sameTimeInBeijing,
			},
			[]byte(`2009-02-01T23:04:05+0800 Beijing`),
			nil,
			false,
		},
		{
			"big number float",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 42.1e+15,
				},
			},
			[]byte(`4.21e+16`),
			nil,
			false,
		},
		{
			"infinity",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: math.MaxFloat32,
				},
			},
			[]byte(`INF`),
			nil,
			false,
		},
		{
			"dict",
			fields{
				kind: HaystackTypeDict,
				dict: map[string]*Value{
					"a": {
						kind: HaystackTypeStr,
						str:  &abc,
					},
				},
			},
			[]byte(`{a:"abc"}`),
			nil,
			false,
		},
		{
			"dict bad element",
			fields{
				kind: HaystackTypeDict,
				dict: map[string]*Value{
					"a": {
						kind: HaystackTypeStr,
						str:  &abc,
					},
					"bad": {
						kind: HaystackTypeUndefined,
					},
				},
			},
			nil,
			nil,
			true,
		},
		{
			"dict two elements",
			fields{
				kind: HaystackTypeDict,
				dict: map[string]*Value{
					"a": {
						kind: HaystackTypeStr,
						str:  &abc,
					},
					"b": {
						kind: HaystackTypeStr,
						str:  &abc,
					},
				},
			},
			[]byte(`{a:"abc",b:"abc"}`),
			[]byte(`{b:"abc",a:"abc"}`),
			false,
		},
		{
			"list one element",
			fields{
				kind: HaystackTypeList,
				list: []*Value{
					{
						kind: HaystackTypeStr,
						str:  &abc,
					},
				},
			},
			[]byte(`["abc"]`),
			nil,
			false,
		},
		{
			"list bad type",
			fields{
				kind: HaystackTypeList,
				list: []*Value{
					{
						kind: HaystackTypeUndefined,
					},
				},
			},
			nil,
			nil,
			true,
		},
		{
			"list",
			fields{
				kind: HaystackTypeList,
				list: []*Value{
					{
						kind: HaystackTypeStr,
						str:  &abc,
					},
					{
						kind: HaystackTypeStr,
						str:  &def,
					},
				},
			},
			[]byte(`["abc","def"]`),
			nil,
			false,
		},

		{
			"-infinity",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: -math.MaxFloat32,
				},
			},
			[]byte(`-INF`),
			nil,
			false,
		},
		{
			"ref",
			fields{
				kind: HaystackTypeRef,
				ref:  NewHaystackID("bla"),
			},
			[]byte(`@bla`),
			nil,
			false,
		},
		{
			"bad grid",
			fields{
				kind: HaystackTypeGrid,
				g:    &Grid{},
			},
			nil,
			nil,
			true,
		},
		{
			"good grid",
			fields{
				kind: HaystackTypeGrid,
				g: &Grid{
					Meta: map[string]string{
						"ver": "3.0",
					},
				},
			},
			[]byte("<<\nver:\"3.0\"\n\n>>\n"),
			nil,
			false,
		},
		{
			"default",
			fields{
				kind: Kind(HaystackLastType + 1),
			},
			nil,
			nil,
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				kind:   tt.fields.kind,
				str:    tt.fields.str,
				number: tt.fields.number,
				b:      tt.fields.b,
				t:      tt.fields.t,
				u:      tt.fields.u,
				ref:    tt.fields.ref,
				g:      tt.fields.g,
				dict:   tt.fields.dict,
				list:   tt.fields.list,
				coord:  tt.fields.coord,
			}
			got, err := v.MarshalZinc()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.MarshalZinc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.orwant == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Value.MarshalZinc() = %v, want %v", string(got), string(tt.want))
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.orwant) {
					t.Errorf("Value.MarshalZinc() = %v, want %v", string(got), string(tt.want))
				}

			}
		})
	}
}
