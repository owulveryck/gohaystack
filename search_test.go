package gohaystack

import (
	"reflect"
	"testing"
)

func TestGrid_GetRowsMatching(t *testing.T) {
	tests := []struct {
		name    string
		fields  *Grid
		args    map[string]*TypedValue
		want    [][]*TypedValue
		wantErr bool
	}{
		{
			name:   "No Match 1",
			fields: &Grid{},
			args: map[string]*TypedValue{
				"col1": {},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "No Match 2",
			fields: &Grid{
				db: map[string][]*TypedValue{
					"col1": nil,
				},
			},
			args: map[string]*TypedValue{
				"col1": nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Simple match",
			fields: &Grid{
				db: map[string][]*TypedValue{
					"col1": {
						{Value: 42},
						{Value: 43},
						{Value: 44},
					},
				},
				Cols: map[int]string{
					0: "col1",
				},
			},
			args: map[string]*TypedValue{
				"col1": {Value: 42},
			},
			want: [][]*TypedValue{
				{
					&TypedValue{Value: 42},
				},
			},
			wantErr: false,
		},
		{
			name: "Simple match, multi rows",
			fields: &Grid{
				db: map[string][]*TypedValue{
					"col1": {
						{Value: 42},
						{Value: 43},
						{Value: 44},
						{Value: 42},
					},
				},
				Cols: map[int]string{
					0: "col1",
				},
			},
			args: map[string]*TypedValue{
				"col1": {Value: 42},
			},
			want: [][]*TypedValue{
				{
					{Value: 42},
				},
				{
					{Value: 42},
				},
			},
			wantErr: false,
		},
		{
			name: "Simple match, multi rows, multi colums",
			fields: &Grid{
				db: map[string][]*TypedValue{
					"col1": {
						{Value: 42},
						{Value: 43},
						{Value: 44},
						{Value: 42},
					},
					"col2": {
						{},
						{},
						{},
						{},
					},
				},
				Cols: map[int]string{
					0: "col1",
					1: "col2",
				},
			},
			args: map[string]*TypedValue{
				"col1": {Value: 42},
			},
			want: [][]*TypedValue{
				{
					{Value: 42},
					{},
				},
				{
					{Value: 42},
					{},
				},
			},
			wantErr: false,
		},
		{
			name: "Multi match, multi rows, multi colums",
			fields: &Grid{
				db: map[string][]*TypedValue{
					"col1": {
						{Value: 42},
						{Value: 43},
						{Value: 44},
						{Value: 42},
					},
					"col2": {
						{Value: 42.42},
						{},
						{},
						{},
					},
				},
				Cols: map[int]string{
					0: "col1",
					1: "col2",
				},
			},
			args: map[string]*TypedValue{
				"col1": {Value: 42},
				"col2": {Value: 42.42},
			},
			want: [][]*TypedValue{
				{
					{Value: 42},
					{Value: 42.42},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Meta:         tt.fields.Meta,
				db:           tt.fields.db,
				colsDis:      tt.fields.colsDis,
				Cols:         tt.fields.Cols,
				lastCol:      tt.fields.lastCol,
				numberOfRows: tt.fields.numberOfRows,
			}
			got, err := g.GetRowsMatching(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grid.GetRowsMatching() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("Grid.GetRowsMatching() = len(%v), want len(%v)", got, tt.want)
			}
			for i := 0; i < len(got); i++ {
				if len(got[i]) != len(tt.want[i]) {
					t.Errorf("Grid.GetRowsMatching() = len(%v), want len(%v)", got[i], tt.want[i])
				}
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("Grid.GetRowsMatching() = %v, want %v", got[i], tt.want[i])
				}
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] == tt.want[i][j] && got[i][j] == nil {
						continue
					}
					if got[i][j].Type != tt.want[i][j].Type ||
						got[i][j].Value != tt.want[i][j].Value {
						t.Errorf("Grid.GetRowsMatching() = %v, want %v", got[i][j], tt.want[i][j])

					}
				}
			}
		})
	}
}

func Test_inter(t *testing.T) {
	type args struct {
		arrs [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"empty arrray",
			args{
				[][]int{
					{1, 2, 3, 4},
					{},
				},
			},
			[]int{},
		},
		{
			"test1",
			args{
				[][]int{
					{1, 2, 3, 4},
					{2, 3, 4},
				},
			},
			[]int{2, 3, 4},
		},
		{
			"test2",
			args{
				[][]int{
					{1, 2, 3, 4},
					{2, 3, 4},
					{3, 4},
					{5, 4},
				},
			},
			[]int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inter(tt.args.arrs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inter() = %v, want %v", got, tt.want)
			}
		})
	}
}
