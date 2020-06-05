package gohaystack

import (
	"reflect"
	"testing"
)

func TestGrid_UnmarshalJSON(t *testing.T) {
	grid := NewGrid()
	err := grid.UnmarshalJSON(samplePayload)
	if err != nil {
		t.Fatal(err)
	}
	if grid.numberOfRows != 3 {
		t.Fail()
	}
}

func TestGrid_NewRow(t *testing.T) {
	type fields struct {
		Meta         map[string]string
		db           map[string][]*TypedValue
		colsDis      map[string]string
		Cols         map[int]string
		lastCol      int
		numberOfRows int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "add row",
			fields: fields{
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
				numberOfRows: 1,
			},
			want: 1,
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
			if got := g.NewRow(); got != tt.want {
				t.Errorf("Grid.NewRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_CloneStruct(t *testing.T) {
	testGrid := NewGrid()
	testGrid.AddColumn("col1", "la colonne 1 (string)")
	testGrid.AddColumn("col2", "")
	testGrid.AddColumn("col3", "array")
	testGrid.AddColumn("col4", "")
	err := testGrid.AddRow([]*TypedValue{
		NewTypedValue(HaystackTypeStr, "bla"),
		NewTypedValue(HaystackTypeStr, "blo"),
		NewTypedValue(HaystackTypeStr, "blu"),
		NewTypedValue(HaystackTypeStr, "bli"),
	})
	if err != nil {
		t.Fatal(err)
	}
	newGrid := testGrid.CloneStruct()
	if !reflect.DeepEqual(newGrid.colsDis, testGrid.colsDis) {
		t.Fail()
	}
	if !reflect.DeepEqual(newGrid.Cols, testGrid.Cols) {
		t.Fail()
	}
	if !reflect.DeepEqual(newGrid.lastCol, testGrid.lastCol) {
		t.Fail()
	}
	if !reflect.DeepEqual(newGrid.Meta, testGrid.Meta) {
		t.Fail()
	}
}

func TestGrid_GetCol(t *testing.T) {
	type fields struct {
		Meta         map[string]string
		db           map[string][]*TypedValue
		colsDis      map[string]string
		Cols         map[int]string
		lastCol      int
		numberOfRows int
	}
	type args struct {
		col string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*TypedValue
		want1  bool
	}{
		{
			"ok",
			fields{
				db: map[string][]*TypedValue{"col1": []*TypedValue{nil}},
			},
			args{
				"col1",
			},
			[]*TypedValue{nil},
			true,
		},
		{
			"ko",
			fields{
				db: map[string][]*TypedValue{"col1": []*TypedValue{nil}},
			},
			args{
				"col2",
			},
			nil,
			false,
		},
		// TODO: Add test cases.
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
			got, got1 := g.GetCol(tt.args.col)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.GetCol() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Grid.GetCol() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
