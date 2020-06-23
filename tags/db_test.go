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
