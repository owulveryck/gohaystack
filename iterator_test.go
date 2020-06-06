package gohaystack

import (
	"testing"
)

func TestNewRowIterator(t *testing.T) {
	g := NewGrid()
	g.AddColumn("col1", "la colonne 1 (string)")
	g.AddColumn("col2", "")
	g.AddColumn("col3", "string")
	for i := 0; i < 5; i++ {
		err := g.AddRow([]*TypedValue{
			NewTypedValue(HaystackTypeNumber, i),
			NewTypedValue(HaystackTypeStr, "bli"),
			NewTypedValue(HaystackTypeStr, "ble"),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	it := NewRowIterator(g)
	if it.Len() != 5 {
		t.Fail()
	}
	for i := 0; it.Next(); i++ {
		if it.Len() != 5-i-1 {
			t.Fail()
		}
		row := it.Row()
		if row[0].Value.(int) != i {
			t.Fail()

		}
	}
}
