package tags

import (
	"reflect"
	"testing"

	"github.com/owulveryck/gohaystack"
)

func TestMarker(t *testing.T) {
	label := gohaystack.NewLabel("bla")
	type args struct {
		l *gohaystack.Label
	}
	tests := []struct {
		name string
		args args
		want func() (*gohaystack.Label, *gohaystack.Value)
	}{
		{
			"name",
			args{
				l: label,
			},
			func() (*gohaystack.Label, *gohaystack.Value) {
				return label, gohaystack.MarkerValue
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Marker(tt.args.l)
			gotLabel, gotValue := got()
			expectedLabel, expectedValue := tt.want()
			if !reflect.DeepEqual(gotLabel, expectedLabel) {
				t.Errorf("Marker() error")
			}
			if !reflect.DeepEqual(gotValue, expectedValue) {
				t.Errorf("Marker() error")
			}
		})
	}
}
