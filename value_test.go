package gohaystack

import (
	"math"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestValue_MarshalJSON(t *testing.T) {
	abc := "abc"
	def := "def"
	simpleTestStr := "test"
	id := NewHaystackID("id")
	u, _ := url.Parse("https://example.com")
	kwh := NewUnit("kwh")
	//simpleTestStrWithColon := "test:bla"
	type fields struct {
		kind   Kind
		str    *string
		number struct {
			value float32
			unit  Unit
		}
		dict  map[string]*Value
		list  []*Value
		b     bool
		t     time.Time
		u     *url.URL
		ref   *HaystackID
		g     *Grid
		coord struct {
			long float32
			lat  float32
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"simple url",
			fields{
				kind: HaystackTypeURI,
				u:    u,
			},
			[]byte(`"u:https://example.com"`),
			false,
		},
		{
			"simple marker",
			fields{
				kind: HaystackTypeMarker,
			},
			[]byte(`"m:"`),
			false,
		},
		{
			"simple reference",
			fields{
				kind: HaystackTypeRef,
				ref:  id,
			},
			[]byte(`"r:` + string(*id) + `"`),
			false,
		},
		{
			"simple number without unit",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 32.0,
				},
			},
			[]byte(`"n:32"`),
			false,
		},
		{
			"simple number with unit",
			fields{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 32.0,
					unit:  kwh,
				},
			},
			[]byte(`"n:32 kwh"`),
			false,
		},
		{
			"simple string",
			fields{
				kind: HaystackTypeStr,
				str:  &simpleTestStr,
			},
			[]byte(`"s:` + simpleTestStr + `"`),
			false,
		},
		{
			"boolean",
			fields{
				kind: HaystackTypeBool,
				b:    true,
			},
			[]byte(`true`),
			false,
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
			[]byte(`["s:abc","s:def"]`),
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
			[]byte(`{"a":"s:abc"}`),
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
					lat:  37.545,
					long: -77.449,
				},
			},
			[]byte(`"c:37.545,-77.449"`),
			false,
		},
		{
			"date",
			fields{
				kind: HaystackTypeDate,
				t:    time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			[]byte(`"d:2009-11-10"`),
			false,
		},
		{
			"time",
			fields{
				kind: HaystackTypeTime,
				t:    time.Date(2009, time.November, 10, 23, 01, 02, 0, time.UTC),
			},
			[]byte(`"h:23:01:02"`),
			false,
		},
		{
			"date time",
			fields{
				kind: HaystackTypeDateTime,
				t:    time.Date(2009, time.November, 10, 23, 01, 02, 0, time.UTC),
			},
			[]byte(`"t:2009-11-10T23:01:02Z"`),
			false,
		},
		{
			"remove",
			fields{
				kind: HaystackTypeRemove,
			},
			[]byte(`"-:"`),
			false,
		},
		{
			"grid",
			fields{
				kind: HaystackTypeGrid,
				g:    NewGrid(),
			},
			[]byte(`{"meta":{"ver":"3.0"}}`),
			false,
		},
		{
			"na",
			fields{
				kind: HaystackTypeNA,
			},
			[]byte(`"z:"`),
			false,
		},
		{
			"Unhandled",
			fields{
				kind: HaystackLastType,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				kind:   tt.fields.kind,
				str:    tt.fields.str,
				dict:   tt.fields.dict,
				list:   tt.fields.list,
				number: tt.fields.number,
				b:      tt.fields.b,
				t:      tt.fields.t,
				u:      tt.fields.u,
				ref:    tt.fields.ref,
				g:      tt.fields.g,
				coord:  tt.fields.coord,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestValue_GetString(t *testing.T) {
	testStr := "testStr"
	type fields struct {
		kind   Kind
		str    *string
		number struct {
			value float32
			unit  Unit
		}
		dict  map[string]*Value
		list  []*Value
		b     bool
		t     time.Time
		u     *url.URL
		ref   *HaystackID
		g     *Grid
		coord struct {
			long float32
			lat  float32
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"string",
			fields{
				kind: HaystackTypeStr,
				str:  &testStr,
			},
			testStr,
			false,
		},
		{
			"no string",
			fields{
				kind: HaystackLastType,
				str:  &testStr,
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				kind:   tt.fields.kind,
				str:    tt.fields.str,
				number: tt.fields.number,
				dict:   tt.fields.dict,
				list:   tt.fields.list,
				b:      tt.fields.b,
				t:      tt.fields.t,
				u:      tt.fields.u,
				ref:    tt.fields.ref,
				g:      tt.fields.g,
				coord:  tt.fields.coord,
			}
			got, err := v.GetString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_UnmarshalJSON(t *testing.T) {
	blabla := "blabla"
	dict := `Dict!`
	type fields struct {
		kind   Kind
		str    *string
		number struct {
			value float32
			unit  Unit
		}
		dict  map[string]*Value
		list  []*Value
		b     bool
		t     time.Time
		u     *url.URL
		ref   *HaystackID
		g     *Grid
		coord struct {
			long float32
			lat  float32
		}
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
			"nil value",
			fields{},
			args{
				nil,
			},
			true,
		},
		{
			"empty value",
			fields{},
			args{
				[]byte(``),
			},
			true,
		},
		{
			"string",
			fields{
				kind: HaystackTypeStr,
				str:  &blabla,
			},
			args{
				[]byte(`"s:blabla"`),
			},
			false,
		},
		{
			"list",
			fields{
				kind: HaystackTypeList,
				list: []*Value{
					{
						kind: HaystackTypeNumber,
						number: struct {
							value float32
							unit  Unit
						}{
							value: 1,
						},
					},
					{
						kind: HaystackTypeNumber,
						number: struct {
							value float32
							unit  Unit
						}{
							value: 2,
						},
					},
					{
						kind: HaystackTypeNumber,
						number: struct {
							value float32
							unit  Unit
						}{
							value: 3,
						},
					},
				},
			},
			args{
				[]byte(`["n:1", "n:2", "n:3"]`),
			},
			false,
		},
		{
			"dict",
			fields{
				kind: HaystackTypeDict,
				dict: map[string]*Value{
					"dis": {
						kind: HaystackTypeStr,
						str:  &dict,
					},
					"foo": {
						kind: HaystackTypeMarker,
					},
				},
			},
			args{
				[]byte(`{"dis":"Dict!", "foo":"m:"}`),
			},
			false,
		},
		{
			"boolean",
			fields{
				kind: HaystackTypeBool,
				b:    true,
			},
			args{
				[]byte(`true`),
			},
			false,
		},
		{
			"bad entry boolean",
			fields{
				kind: HaystackTypeBool,
				b:    true,
			},
			args{
				[]byte(`ttttrue`),
			},
			true,
		},
		{
			"grid",
			fields{
				kind: HaystackTypeGrid,
				g: &Grid{
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
			},
			args{
				[]byte(`{"meta":{"ver":"3.0"}, "cols":[{"name": "blabla"}],"rows":[{"id":"r:myid","blabla":"blabla"}]}`),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				kind:   tt.fields.kind,
				str:    tt.fields.str,
				dict:   tt.fields.dict,
				list:   tt.fields.list,
				number: tt.fields.number,
				b:      tt.fields.b,
				t:      tt.fields.t,
				u:      tt.fields.u,
				ref:    tt.fields.ref,
				g:      tt.fields.g,
				coord:  tt.fields.coord,
			}
			if err := v.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Value.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewRef(t *testing.T) {
	id := NewHaystackID("myid")
	type args struct {
		r *HaystackID
	}
	tests := []struct {
		name string
		args args
		want *Value
	}{
		{
			"ref",
			args{
				id,
			},
			&Value{
				kind: HaystackTypeRef,
				ref:  id,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRef(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNumber(t *testing.T) {
	kwh := NewUnit("Kwh")
	type args struct {
		value float32
		unit  Unit
	}
	tests := []struct {
		name string
		args args
		want *Value
	}{
		{
			"simple test",
			args{
				value: 32.0,
				unit:  kwh,
			},
			&Value{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					32.0,
					kwh,
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNumber(tt.args.value, tt.args.unit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewURL(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name string
		args args
		want *Value
	}{
		{
			"simple",
			args{
				u,
			},
			&Value{
				kind: HaystackTypeURI,
				u:    u,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_unmarshalJSONString(t *testing.T) {
	bladibla := `bladibla`
	bladiblablabla := `bladibla blabla`
	bladiblas := `blablas:bladibla blabla`
	bladiblaRef := HaystackID(`bladibla`)
	kwh := `kwh`
	date := time.Date(1970, time.January, 01, 00, 00, 00, 00, time.UTC)
	u, _ := url.Parse("http://bing.com/search?q=dotnet")
	tz := "Asia/Shanghai"
	var testTZData []byte
	testTZ, err := time.LoadLocation(tz)
	if err != nil {
		testTZ = time.UTC
		testTZData = []byte(`t:1970-01-01T00:00:00+00:00 UTC`)
	} else {
		testTZData = []byte(`t:1970-01-01T00:00:00+08:00 ` + tz)
	}
	dateTZ := time.Date(1970, time.January, 01, 00, 00, 00, 00, testTZ)
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
	type args struct {
		b []byte
	}
	tests := []struct {
		name     string
		fields   fields
		expected *Value
		args     args
		wantErr  bool
	}{
		{
			"string without marker",
			fields{},
			&Value{
				kind: HaystackTypeStr,
				str:  &bladibla,
			},
			args{
				[]byte(`bladibla`),
			},
			false,
		},
		{
			"string with marker",
			fields{},
			&Value{
				kind: HaystackTypeStr,
				str:  &bladibla,
			},
			args{
				[]byte(`s:bladibla`),
			},
			false,
		},
		{
			"string with marker",
			fields{},
			&Value{
				kind: HaystackTypeStr,
				str:  &bladiblablabla,
			},
			args{
				[]byte(`s:bladibla blabla`),
			},
			false,
		},
		{
			"string with marker",
			fields{},
			&Value{
				kind: HaystackTypeStr,
				str:  &bladiblas,
			},
			args{
				[]byte(`blablas:bladibla blabla`),
			},
			false,
		},
		{
			"marker",
			fields{},
			&Value{
				kind: HaystackTypeMarker,
			},
			args{
				[]byte(`m:`),
			},
			false,
		},
		{
			"reference",
			fields{},
			&Value{
				kind: HaystackTypeRef,
				ref:  &bladiblaRef,
			},
			args{
				[]byte(`r:bladibla`),
			},
			false,
		},
		{
			"remove",
			fields{},
			&Value{
				kind: HaystackTypeRemove,
			},
			args{
				[]byte(`-:`),
			},
			false,
		},
		{
			"na",
			fields{},
			&Value{
				kind: HaystackTypeNA,
			},
			args{
				[]byte(`z:`),
			},
			false,
		},
		{
			"number invalid",
			fields{},
			&Value{},
			args{
				[]byte(`n:invalid`),
			},
			true,
		},
		{
			"number INF",
			fields{},
			&Value{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: float32(math.Inf(1)),
				},
			},
			args{
				[]byte(`n:INF`),
			},
			false,
		},
		{
			"number -INF",
			fields{},
			&Value{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: float32(math.Inf(-1)),
				},
			},
			args{
				[]byte(`n:-INF`),
			},
			false,
		},
		{
			"number with unit",
			fields{},
			&Value{
				kind: HaystackTypeNumber,
				number: struct {
					value float32
					unit  Unit
				}{
					value: 42.0,
					unit:  &kwh,
				},
			},
			args{
				[]byte(`n:42.0 kwh`),
			},
			false,
		},
		{
			"invalid number",
			fields{},
			&Value{},
			args{
				[]byte(`n:42.0 kwh bla`),
			},
			true,
		},
		{
			"invalid date",
			fields{},
			&Value{},
			args{
				[]byte(`d:invalid`),
			},
			true,
		},
		{
			"valid date",
			fields{},
			&Value{
				kind: HaystackTypeDate,
				t:    date,
			},
			args{
				[]byte(`d:1970-01-01`),
			},
			false,
		},
		{
			"invalid time",
			fields{},
			&Value{},
			args{
				[]byte(`h:INVALID`),
			},
			true,
		},
		{
			"valid time",
			fields{},
			&Value{
				kind: HaystackTypeTime,
				t:    date,
			},
			args{
				[]byte(`h:00:00:00`),
			},
			false,
		},
		{
			"invalid datetime",
			fields{},
			&Value{},
			args{
				[]byte(`t:INVALID`),
			},
			true,
		},
		{
			"valid datetime without tz",
			fields{},
			&Value{
				kind: HaystackTypeDateTime,
				t:    date,
			},
			args{
				[]byte(`t:1970-01-01T00:00:00Z`),
			},
			false,
		},
		{
			"valid datetime with tz",
			fields{},
			&Value{
				kind: HaystackTypeDateTime,
				t:    dateTZ,
			},
			args{
				testTZData,
			},
			false,
		},
		{
			"valid datetime with invalid tz",
			fields{},
			&Value{},
			args{
				[]byte(`t:1970-01-01T00:00:00-00:00 NOWHERE`),
			},
			true,
		},
		{
			"invalid datetime with valid tz",
			fields{},
			&Value{},
			args{
				[]byte(`t:19.070-01-01T00:00:00-00:00 UTC`),
			},
			true,
		},
		{
			"invalid uri",
			fields{},
			&Value{},
			args{
				append([]byte(`u:aa`), byte(0x7f)),
			},
			true,
		},
		{
			"valid uri",
			fields{},
			&Value{
				kind: HaystackTypeURI,
				u:    u,
			},
			args{
				[]byte(`u:http://bing.com/search?q=dotnet`),
			},
			false,
		},
		{
			"invalid coord, missing long",
			fields{},
			&Value{},
			args{
				[]byte(`c:12`),
			},
			true,
		},
		{
			"invalid coord too many args",
			fields{},
			&Value{},
			args{
				[]byte(`c:12,23,34`),
			},
			true,
		},
		{
			"invalid coord not a number",
			fields{},
			&Value{},
			args{
				[]byte(`c:aa,41`),
			},
			true,
		},
		{
			"invalid coord not a number",
			fields{},
			&Value{},
			args{
				[]byte(`c:41,bb`),
			},
			true,
		},
		{
			"valid coord",
			fields{},
			&Value{
				kind: HaystackTypeCoord,
				coord: struct {
					long float32
					lat  float32
				}{
					lat:  41,
					long: 42,
				},
			},
			args{
				[]byte(`c:41,42`),
			},
			false,
		},
		{
			"Xstr (Not implemented)",
			fields{},
			&Value{},
			args{
				[]byte(`x:41,42`),
			},
			true,
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
			if err := v.unmarshalJSONString(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Value.unmarshalJSONString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.expected, v) {
				t.Log(v.t.Location())
				t.Log(tt.expected.t.Location())
				t.Errorf("Value.MarshalJSON() = %v, want %v", v, tt.expected)
			}
		})
	}
}

func TestValue_GetKind(t *testing.T) {
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
		name   string
		fields fields
		want   Kind
	}{
		{
			"empty",
			fields{},
			HaystackTypeUndefined,
		},
		{
			"simple",
			fields{
				kind: HaystackTypeBool,
			},
			HaystackTypeBool,
		},
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
			if got := v.GetKind(); got != tt.want {
				t.Errorf("Value.GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_GetHaystackID(t *testing.T) {
	myid := HaystackID("bla")
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
		want    *HaystackID
		wantErr bool
	}{
		{
			"not a ref",
			fields{
				kind: HaystackLastType,
			},
			nil,
			true,
		},
		{
			"valid ref",
			fields{
				kind: HaystackTypeRef,
				ref:  &myid,
			},
			&myid,
			false,
		},
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
			got, err := v.GetHaystackID()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.GetHaystackID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.GetHaystackID() = %v, want %v", got, tt.want)
			}
		})
	}
}
