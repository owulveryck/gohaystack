package gohaystack

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"
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
	if err != nil {
		t.Fatal(err)
	}

	testGrid2 := NewGrid()
	testGrid2.AddColumn("col1", "la colonne 1 (string)")
	testGrid2.AddColumn("col2", "")
	testGrid2.AddColumn("col3", "string")
	err = testGrid2.AddRow([]*TypedValue{
		NewTypedValue(HaystackTypeStr, "blu"),
		NewTypedValue(HaystackTypeStr, "bli"),
		NewTypedValue(HaystackTypeStr, "ble"),
	})
	if err != nil {
		t.Fatal(err)
	}

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

func TestTypedValue_stringJSON(t *testing.T) {
	uri, _ := url.Parse("https://bladibla")
	type fields struct {
		Type  HaystackType
		Value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Marker",
			fields{
				HaystackTypeMarker,
				true,
			},
			`"m:"`,
		},
		{
			"DateTime",
			fields{
				HaystackTypeDateTime,
				time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			`"t:2009-11-10T23:00:00Z"`,
		},
		{
			"Time",
			fields{
				HaystackTypeTime,
				time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			`"h:23:00:00"`,
		},
		{
			"Date",
			fields{
				HaystackTypeDate,
				time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			`"d:2009-11-10"`,
		},
		{
			"URI",
			fields{
				HaystackTypeURI,
				uri,
			},
			`"u:https://bladibla"`,
		},
		{
			"Number",
			fields{
				HaystackTypeNumber,
				&HaystackNumber{
					Value: 22.5,
					Unit:  "FF",
				},
			},
			`"n:22.5 FF"`,
		},
		{
			"Number Without unit",
			fields{
				HaystackTypeNumber,
				&HaystackNumber{
					Value: 22.5,
				},
			},
			`"n:22.5"`,
		},
		{
			"Reference",
			fields{
				HaystackTypeRef,
				"blabla",
			},
			`"r:blabla"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &TypedValue{
				Type:  tt.fields.Type,
				Value: tt.fields.Value,
			}
			if got := v.stringJSON(); got != tt.want {
				t.Errorf("TypedValue.stringJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
