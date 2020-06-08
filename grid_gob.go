package gohaystack

import (
	"bytes"
	"encoding/gob"
	"net/url"
)

func init() {
	gob.Register(new(url.URL))
}

type typedValue struct {
	TV    TypedValue
	IsNil bool
}
type pivot struct {
	Meta         map[string]string
	DB           map[string][]typedValue // column based db
	ColsDis      map[string]string       // description of the column
	Cols         map[int]string          // this is a reverse index string is the col name, and int its index
	LastCol      int                     // index of the last column
	NumberOfRows int
}

// GobEncode the grid in a binary form for export
func (g *Grid) GobEncode() ([]byte, error) {
	db := make(map[string][]typedValue, len(g.db))
	for k, v := range g.db {
		db[k] = make([]typedValue, len(v))
		for i := 0; i < len(v); i++ {
			if v[i] == nil {
				db[k][i] = typedValue{
					TV:    TypedValue{},
					IsNil: true,
				}
			} else {
				db[k][i] = typedValue{
					TV:    *v[i],
					IsNil: false,
				}
			}
		}
	}
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(pivot{
		Meta:         g.Meta,
		DB:           db,
		ColsDis:      g.colsDis,
		Cols:         g.Cols,
		LastCol:      g.lastCol,
		NumberOfRows: g.numberOfRows,
	})
	return b.Bytes(), err
}

// GobDecode ...
func (g *Grid) GobDecode(b []byte) error {
	rdr := bytes.NewBuffer(b)
	dec := gob.NewDecoder(rdr)
	var p pivot
	err := dec.Decode(&p)
	if err != nil {
		return err
	}
	g.Meta = p.Meta
	g.colsDis = p.ColsDis
	g.Cols = p.Cols
	g.lastCol = p.LastCol
	g.numberOfRows = p.NumberOfRows
	g.db = make(map[string][]*TypedValue, len(p.DB))
	for k, v := range p.DB {
		g.db[k] = make([]*TypedValue, len(v))
		for i := 0; i < len(v); i++ {
			if !v[i].IsNil {
				g.db[k][i] = &v[i].TV
			}
		}
	}
	return nil
}
