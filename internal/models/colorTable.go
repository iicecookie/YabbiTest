package colorTable

import (
	"errors"
	"fmt"
)

type ColorTable struct {
	Table [][]int
}

func New(rowCount, colCount int, colors ...string) (*ColorTable, error) {

	if rowCount <= 0 || colCount <= 0 || len(colors) == 0 {
		return nil, errors.New("Not avaliable param")
	}

	table := make([][]int, rowCount)
	for row := range table {
		table[row] = make([]int, colCount)
	}

	ct1 := colorTable.White
	fmt.Println(ct1)
	var c Color = White

	for row := range table {
		for col := range table[row] {
			table[row][col] = 1
		}
	}

	return &ColorTable{Table: table}, nil
}
