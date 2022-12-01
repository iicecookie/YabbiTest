package colorTable

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/TwiN/go-color"
)

type ColorTable struct {
	Table [][]Color
}

func New(rowCount, colCount int, colors ...string) (*ColorTable, error) {

	if rowCount <= 0 || colCount <= 0 {
		return nil, errors.New("not avaliable table size value param")
	}
	if len(colors) == 0 {
		return nil, errors.New("there are must be at least one color value")
	}

	// initialize table
	table := make([][]Color, rowCount)
	for row := range table {
		table[row] = make([]Color, colCount)
	}

	err := fillTableRandomColors(table, colors...)
	if err != nil {
		return nil, err
	}

	return &ColorTable{Table: table}, nil
}

func fillTableRandomColors(table [][]Color, colors ...string) error {
	rand.Seed(time.Now().UnixNano())

	const chanceFillCell = 50 // %

	for row := range table {
		for col := range table[row] {

			if rand.Intn(100) > chanceFillCell {

				randColor := rand.Intn(len(colors))
				color, ok := ColorMap[colors[randColor]]
				if !ok {
					return fmt.Errorf("no such color found %s", colors[randColor])
				}
				table[row][col] = color
			}
		}
	}
	return nil
}

func (ct *ColorTable) GetMaxNeighboringColors() int {

	rowCount := len(ct.Table)
	colCount := len(ct.Table[0])

	markTable := make([][]Color, rowCount)
	for row := 0; row < rowCount; row++ {
		markTable[row] = make([]Color, colCount)
	}

	maxColors := 0
	const emptyColor = 0

	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {

			foundColor := ct.Table[row][col]
			if foundColor != emptyColor {

				colorsFound := searchForNeigborColors(ct.Table, markTable, foundColor, row, col, rowCount, colCount)
				if colorsFound > maxColors {
					maxColors = colorsFound
				}
			}
		}
	}
	return maxColors
}

func searchForNeigborColors(table [][]Color, markTable [][]Color, searchingColor Color, row, col, rowCount, colCount int) int {

	if row < 0 || col < 0 || row >= rowCount || col >= colCount {
		return 0
	}

	if table[row][col] != searchingColor || markTable[row][col] != 0 {
		return 0
	}

	markTable[row][col] = 1

	return 1 +
		searchForNeigborColors(table, markTable, searchingColor, row+1, col, rowCount, colCount) +
		searchForNeigborColors(table, markTable, searchingColor, row-1, col, rowCount, colCount) +
		searchForNeigborColors(table, markTable, searchingColor, row, col+1, rowCount, colCount) +
		searchForNeigborColors(table, markTable, searchingColor, row, col-1, rowCount, colCount)

}

func (ct *ColorTable) String() string {

	var result string
	const cellPlaceHolder = "◼"

	for row := 0; row < len(ct.Table); row++ {
		for column := 0; column < len(ct.Table[row]); column++ {

			cellColor := Color(ct.Table[row][column])
			result += " " + color.Ize(cellColor.GetCode(), cellPlaceHolder)
		}
		result += "\n"
	}
	return result
}
