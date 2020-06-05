package gohaystack

import (
	"math"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func Test_inferType(t *testing.T) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}
	refTime := time.Date(2006, time.January, 01, 23, 2, 0, 0, location)
	date, _ := time.Parse("2006-01-02", refTime.Format("2006-01-02"))
	curTime, _ := time.Parse("15:04:00", refTime.Format("15:04:00"))

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    HaystackType
		want1   interface{}
		wantErr bool
	}{
		{
			"Time",
			args{
				value: "h:" + refTime.Format("15:04:00"),
			},
			HaystackTypeTime,
			curTime,
			false,
		},
		{
			"DateTime",
			args{
				value: "t:2006-01-01T23:02:00-05:00 America/New_York",
			},
			HaystackTypeDateTime,
			refTime,
			false,
		},
		{
			"DateTime without TZ",
			args{
				value: "t:2006-01-01T23:02:00-05:00",
			},
			HaystackTypeDateTime,
			refTime,
			false,
		},
		{
			"Date",
			args{
				value: "d:" + refTime.Format("2006-01-02"),
			},
			HaystackTypeDate,
			date,
			false,
		},
		{
			"URI valid",
			args{
				value: "u:s3://blabla/bla",
			},
			HaystackTypeURI,
			&url.URL{
				Scheme: "s3",
				Host:   "blabla",
				Path:   "/bla",
			},
			false,
		},
		{
			"number INF",
			args{
				value: "n:INF °F",
			},
			HaystackTypeNumber,
			&HaystackNumber{
				Unit:  "°F",
				Value: float32(math.Inf(1)),
			},
			false,
		},
		{
			"number -INF",
			args{
				value: "n:-INF °F",
			},
			HaystackTypeNumber,
			&HaystackNumber{
				Unit:  "°F",
				Value: float32(math.Inf(-1)),
			},
			false,
		},
		{
			"number with unit",
			args{
				value: "n:1234 °F",
			},
			HaystackTypeNumber,
			&HaystackNumber{
				Unit:  "°F",
				Value: 1234,
			},
			false,
		},
		{
			"number without unit",
			args{
				value: "n:1234",
			},
			HaystackTypeNumber,
			&HaystackNumber{
				Unit:  "",
				Value: 1234,
			},
			false,
		},

		{
			"number bad format",
			args{
				value: "n:abc",
			},
			HaystackTypeUndefined,
			nil,
			true,
		},
		{
			"marker",
			args{
				value: "m:",
			},
			HaystackTypeMarker,
			true,
			false,
		},
		{
			"String 1",
			args{
				value: "s:blabla",
			},
			HaystackTypeStr,
			"blabla",
			false,
		},
		{
			"String 2",
			args{
				value: "blabla",
			},
			HaystackTypeStr,
			"blabla",
			false,
		},
		{
			"Reference",
			args{
				value: "r:blabla",
			},
			HaystackTypeRef,
			"blabla",
			false,
		},
		{
			"Boolean",
			args{
				value: true,
			},
			HaystackTypeBool,
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, val, err := inferType(tt.args.value)
			if typ != tt.want {
				t.Errorf("inferType() got = %v, want %v", typ, tt.want)
			}
			if typ == HaystackTypeDateTime {
				if !val.(time.Time).Equal(tt.want1.(time.Time)) {
					t.Errorf("inferType() time error got1 = %v, want %v", val, tt.want1)
				}
			} else {
				if !reflect.DeepEqual(val, tt.want1) {
					t.Errorf("inferType() got1 = %v, want %v", val, tt.want1)
				}
			}
			if err == nil && tt.wantErr {
				t.Fail()
			}
		})
	}
}
