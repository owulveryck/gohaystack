package gohaystack

// RowIterator is a top level structure to get the row of a db
type RowIterator struct {
	cur  int
	grid *Grid
}

// NewRowIterator from a grid
func NewRowIterator(g *Grid) *RowIterator {
	return &RowIterator{
		cur:  -1,
		grid: g,
	}
}

// Len returns the remaining number of rows to be iterated over.
func (r *RowIterator) Len() int {
	return r.grid.numberOfRows - r.cur - 1
}

// Row returns the current row of the iterator. Next must have been called prior to a call to Row.
func (r *RowIterator) Row() []*Tag {
	row := make([]*Tag, len(r.grid.Cols))
	for i := 0; i < len(r.grid.Cols); i++ {
		row[i] = r.grid.db[r.grid.Cols[i]][r.cur]
	}
	return row
}

// Next returns whether the next call of Row will return a valid row.
func (r *RowIterator) Next() bool {
	if r.cur == r.grid.numberOfRows {
		return false
	}
	r.cur++
	return r.cur < r.grid.numberOfRows
}
