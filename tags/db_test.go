package tags

import (
	"reflect"
	"testing"

	"github.com/owulveryck/gohaystack"
)

func TestGetLabelByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  *gohaystack.Label
		want1 bool
	}{
		{
			"found",
			args{
				"site",
			},
			labelDB["site"],
			true,
		},
		{
			"not found",
			args{
				"NOT FOUND",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetLabelByName(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLabelByName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLabelByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetMarkerByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  func() (*gohaystack.Label, *gohaystack.Value)
		want1 bool
	}{
		{
			"found",
			args{
				"site",
			},
			markerDB["site"],
			true,
		},
		{
			"not found",
			args{
				"NOTFOUND",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetMarkerByName(tt.args.name)
			if got != nil && tt.want != nil {
				gotL, gotV := got()
				wantL, wantV := tt.want()
				if !reflect.DeepEqual(gotL, wantL) {
					t.Errorf("GetMarkerByName()")
				}
				if !reflect.DeepEqual(gotV, wantV) {
					t.Errorf("GetMarkerByName()")
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetMarkerByName()")
				}
			}
			if got1 != tt.want1 {
				t.Errorf("GetMarkerByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
