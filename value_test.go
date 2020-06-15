package gohaystack

import (
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
			[]byte(`{"meta":{"Ver":"3.0"},"cols":[],"rows":[]}`),
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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"string without marker",
			fields{},
			args{
				[]byte(`bladibla`),
			},
			false,
		},
		{
			"string with marker",
			fields{},
			args{
				[]byte(`s:bladibla`),
			},
			false,
		},
		{
			"string with marker",
			fields{},
			args{
				[]byte(`s:bladibla blabla`),
			},
			false,
		},
		{
			"string with marker",
			fields{},
			args{
				[]byte(`blablas:bladibla blabla`),
			},
			false,
		},
		{
			"reference",
			fields{},
			args{
				[]byte(`r:bladibla`),
			},
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
			if err := v.unmarshalJSONString(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Value.unmarshalJSONString() error = %v, wantErr %v", err, tt.wantErr)
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
