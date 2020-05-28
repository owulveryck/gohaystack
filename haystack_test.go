package gohaystack

import (
	"encoding/json"
	"testing"
)

func Test_grid_MarshalJSON(t *testing.T) {
	testGrid := NewGrid()
	testGrid.AddColumn("col1", "la colonne 1 (string)")
	testGrid.AddColumn("col2", "")
	err := testGrid.AddRow([]*TypedValue{
		NewTypedValue(HaystackTypeStr, "bla"),
		NewTypedValue(HaystackTypeStr, "blo"),
		NewTypedValue(HaystackTypeStr, "blu"),
		NewTypedValue(HaystackTypeStr, "bli"),
	})
	if err == nil {
		t.Fatal(err)
	}
	testGrid.AddColumn("col3", "array")
	testGrid.AddColumn("col4", "")
	testGrid.AddColumn("col5", "grid")
	err = testGrid.AddRow([]*TypedValue{
		NewTypedValue(HaystackTypeStr, "bla"),
		NewTypedValue(HaystackTypeStr, "blo"),
		NewTypedValue(HaystackTypeStr, "blu"),
		NewTypedValue(HaystackTypeStr, "bli"),
		NewTypedValue(HaystackTypeStr, "ble"),
	})

	testGrid2 := NewGrid()
	testGrid2.AddColumn("col1", "la colonne 1 (string)")
	testGrid2.AddColumn("col2", "")
	testGrid2.AddColumn("col3", "string")
	err = testGrid2.AddRow([]*TypedValue{
		NewTypedValue(HaystackTypeStr, "blu"),
		NewTypedValue(HaystackTypeStr, "bli"),
		NewTypedValue(HaystackTypeStr, "ble"),
	})

	err = testGrid.AddRow([]*TypedValue{
		NewTypedValue(HaystackTypeStr, "bla2"),
		NewTypedValue(HaystackTypeStr, "blo2"),
		NewTypedValue(HaystackTypeStr, "blu2"),
		NewTypedValue(HaystackTypeStr, "bli2"),
		NewTypedValue(HaystackTypeGrid, testGrid2),
	})
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.MarshalIndent(testGrid, "\t", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func Test_haystackRow_MarshalJSON(t *testing.T) {
	hr := haystackRow{
		haystackKVPair{
			Name: "test",
			Value: &TypedValue{
				Value: "a",
			},
		},
		haystackKVPair{
			Name: "test2",
			Value: &TypedValue{
				Value: "b",
			},
		},
	}
	b, err := json.MarshalIndent(hr, "\t", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
