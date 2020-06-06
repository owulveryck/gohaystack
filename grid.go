package gohaystack

import (
	"encoding/json"
	"errors"
)

// NewGrid creates an empty grid
func NewGrid() *Grid {
	return &Grid{
		Meta: map[string]string{
			"Ver": "3,0",
		},
		db:      make(map[string][]*TypedValue),
		colsDis: make(map[string]string),
		Cols:    make(map[int]string),
	}
}

// Grid is a simple database structure, column based
type Grid struct {
	Meta         map[string]string
	db           map[string][]*TypedValue // column based db
	colsDis      map[string]string        // description of the column
	Cols         map[int]string           // this is a reverse index string is the col name, and int its index
	lastCol      int                      // index of the last column
	numberOfRows int
}

// CloneStruct ...
func (g *Grid) CloneStruct() *Grid {
	colsNum := len(g.db)
	meta := make(map[string]string, len(g.Meta))
	colsDis := make(map[string]string, colsNum)
	cols := make(map[int]string, colsNum)
	for k, v := range g.Meta {
		meta[k] = v
	}
	for k, v := range g.colsDis {
		colsDis[k] = v
	}
	for k, v := range g.Cols {
		cols[k] = v
	}
	emptyDB := make(map[string][]*TypedValue, colsNum)
	for k := range g.db {
		emptyDB[k] = nil
	}

	return &Grid{
		Meta:    meta,
		db:      emptyDB,
		colsDis: colsDis,
		Cols:    cols,
		lastCol: g.lastCol,
	}
}

// AddColumn with name and description. This column does not extend the existing rows and therefore is not safe
// for use after AddRow has been called.
// Column is added at the end of the stack, order of call to this function matters.
func (g *Grid) AddColumn(name, dis string) {
	g.db[name] = make([]*TypedValue, g.numberOfRows)
	g.Cols[g.lastCol] = name
	g.colsDis[name] = dis
	g.lastCol++
}

// NewRow empty row in the grid; returns the rowID
func (g *Grid) NewRow() int {
	g.numberOfRows++
	for k := range g.db {
		g.db[k] = append(g.db[k], nil)
	}
	return g.numberOfRows - 1
}

// GetCol returns the column value and a boolean to false if the column does not exists
func (g *Grid) GetCol(col string) ([]*TypedValue, bool) {
	res, ok := g.db[col]
	return res, ok
}

// Set the value in the grid; silently override any existing value.
// returns an error if col or row is out of range.
func (g *Grid) Set(row int, col string, value *TypedValue) error {
	if _, ok := g.db[col]; !ok {
		return errors.New("Non existent column " + col)
	}
	if row > g.numberOfRows-1 {
		return errors.New("bad row number")
	}
	g.db[col][row] = value
	return nil
}

// AddRow to the grid at the end of the existing row stack
func (g *Grid) AddRow(row []*TypedValue) error {
	if len(row) != g.lastCol {
		return errors.New("row size does not fit number of cols")
	}
	for i := 0; i < g.lastCol; i++ {
		g.db[g.Cols[i]] = append(g.db[g.Cols[i]], row[i])
	}
	g.numberOfRows++
	return nil
}

// MarshalJSON encode the grid in a haystack compatible format
func (g *Grid) MarshalJSON() ([]byte, error) {
	var h haystack
	h.Meta = g.Meta
	h.Cols = make([]haystackCol, g.lastCol)
	h.Rows = make([]haystackRow, g.numberOfRows)
	for i := 0; i < g.lastCol; i++ {
		h.Cols[i] = haystackCol{
			Name: g.Cols[i],
			Dis:  g.colsDis[g.Cols[i]],
		}
	}
	for i := 0; i < g.numberOfRows; i++ {
		row := make([]haystackKVPair, g.lastCol)
		for j := 0; j < g.lastCol; j++ {
			row[j] = haystackKVPair{
				Name:  g.Cols[j],
				Value: g.db[g.Cols[j]][i],
			}
		}
		h.Rows[i] = row
	}
	return json.Marshal(h)
}

// UnmarshalJSON implemented the json.Unmarshaler. Reads a haystack json payload and creates a grid.
func (g *Grid) UnmarshalJSON(b []byte) error {
	type haystackStructure struct {
		Meta struct {
			Ver string `json:"ver"`
		} `json:"meta"`
		Cols []struct {
			Name string `json:"name"`
			Dis  string `json:"dis"`
		} `json:"cols"`
		Rows []map[string]interface{} `json:"rows"`
	}
	var temp haystackStructure
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	for i := range temp.Cols {
		g.AddColumn(temp.Cols[i].Name, temp.Cols[i].Dis)
	}
	for i := range temp.Rows {
		r := g.NewRow()
		for k, v := range temp.Rows[i] {
			typ, val, err := inferType(v)
			if err != nil {
				return err
			}
			g.Set(r, k, NewTypedValue(typ, val))
		}
	}
	return nil
}
