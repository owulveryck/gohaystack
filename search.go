package gohaystack

import (
	"errors"
	"sort"
)

// GetRowsMatching returns rows matching the match hashmap.
// The key of match is the column name and *TypedValue a non nil value to be matched.
// The return value is a list of rows (the row is of length len(g.Cols))
// This function returns an error if the key of match does not exist in the grid
// and if any of the pointer to TypedValue is nil in the request.
func (g *Grid) GetRowsMatching(match map[string]*TypedValue) ([][]*TypedValue, error) {
	allMatchingRows := make([][]int, 0)
	for k, v := range match {
		if v == nil {
			return nil, errors.New("Cannot search nil elements")
		}
		var col []*TypedValue
		var ok bool
		if col, ok = g.db[k]; !ok {
			return nil, errors.New("Column does not exist " + k)
		}
		matchingRowsNumbers := make([]int, 0)
		for i := 0; i < len(col); i++ {
			if col[i].Equal(v) {
				matchingRowsNumbers = append(matchingRowsNumbers, i)
			}
		}
		// No match!
		if len(matchingRowsNumbers) == 0 {
			return nil, nil
		}
		allMatchingRows = append(allMatchingRows, matchingRowsNumbers)
	}
	var rowsID []int
	if len(allMatchingRows) == 1 {
		rowsID = allMatchingRows[0]
	} else {
		rowsID = inter(allMatchingRows...)
	}
	rows := make([][]*TypedValue, len(rowsID))
	for i := 0; i < len(rows); i++ {
		id := rowsID[i]
		rows[i] = make([]*TypedValue, len(g.Cols))
		for c := 0; c < len(g.Cols); c++ {
			column := g.Cols[c]
			rows[i][c] = g.db[column][id]
		}
	}
	return rows, nil
}

func inter(arrs ...[]int) []int {
	for i := 0; i < len(arrs); i++ {
		if len(arrs[i]) == 0 {
			return []int{}
		}
		sort.Ints(arrs[i])
	}
	res := []int{}
	x := arrs[0][0]
	i := 1
	for {
		off := sort.SearchInts(arrs[i], x)
		if off == len(arrs[i]) {
			// we emptied one slice, we're done.
			break
		}
		if arrs[i][off] == x {
			i++
			if i == len(arrs) {
				// x was in all the slices
				res = append(res, x)
				x++ // search for the next possible x.
				i = 0
			}
		} else {
			x = arrs[i][off]
			i = 0 // This can be done a bit more optimally.
		}
	}
	return res
}
