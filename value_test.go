package gohaystack

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestValue_MarshalJSON(t *testing.T) {
	simpleTestStr := "test"
	id := NewHaystackID("id")
	u, _ := url.Parse("https://example.com")
	//simpleTestStrWithColon := "test:bla"
	type fields struct {
		kind   kind
		str    *string
		number struct {
			value float32
			unit  string
		}
		t     *time.Time
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
				kind: haystackTypeURI,
				u:    u,
			},
			[]byte(`"u:https://example.com"`),
			false,
		},
		{
			"simple marker",
			fields{
				kind: haystackTypeMarker,
			},
			[]byte(`"m:"`),
			false,
		},
		{
			"simple reference",
			fields{
				kind: haystackTypeRef,
				ref:  id,
			},
			[]byte(`"r:` + string(*id) + `"`),
			false,
		},
		{
			"simple string",
			fields{
				kind: haystackTypeStr,
				str:  &simpleTestStr,
			},
			[]byte(`"s:` + simpleTestStr + `"`),
			false,
		},
		{
			"Unhandled",
			fields{
				kind: haystackLastType,
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
				number: tt.fields.number,
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
				t.Errorf("Value.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_GetString(t *testing.T) {
	testStr := "testStr"
	type fields struct {
		kind   kind
		str    *string
		number struct {
			value float32
			unit  string
		}
		t     *time.Time
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
				kind: haystackTypeStr,
				str:  &testStr,
			},
			testStr,
			false,
		},
		{
			"no string",
			fields{
				kind: haystackLastType,
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
	type fields struct {
		kind   kind
		str    *string
		number struct {
			value float32
			unit  string
		}
		t     *time.Time
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
		/*
			{
				"string",
				fields{},
				args{
					[]byte(`"s:blabla"`),
				},
				false,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				kind:   tt.fields.kind,
				str:    tt.fields.str,
				number: tt.fields.number,
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
				kind: haystackTypeRef,
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
