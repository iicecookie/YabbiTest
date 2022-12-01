package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/iicecookie/YabbiTest/internal/models/colorTable"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("App Started")

	for {

		fmt.Println("Enter row and col count: ")

		rowCount, colCount := ReadRowAndColCount(scanner)

		avaliableColors := make([]string, 0, len(colorTable.ColorMap))
		for k := range colorTable.ColorMap {
			avaliableColors = append(avaliableColors, k)
		}
		fmt.Println("Enter colors separated by a space. Avaliable colors: ", avaliableColors)

		colors := ReadTableColors(scanner)

		table, err := colorTable.New(rowCount, colCount, colors...)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(table)
		fmt.Println("Maximum neighborhood color sequence", table.GetMaxNeighboringColors())

		fmt.Println("Continue? (y/n)")
		scanner.Scan()
		scanedLine := scanner.Text()
		if scanedLine == "n" {
			break
		}

	}
	fmt.Println("App closed")
}

func ReadRowAndColCount(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	scanedLine := scanner.Text()
	parsedArr := strings.Fields(scanedLine)

	rowCount, err := strconv.Atoi(parsedArr[0])
	if err != nil {
		log.Fatal("error while parsing input row value ")
	}

	colCount, err := strconv.Atoi(parsedArr[1])
	if err != nil {
		log.Fatal("error while parsing input col value ")
	}
	return rowCount, colCount
}

func ReadTableColors(scanner *bufio.Scanner) []string {

	scanner.Scan()
	scanedLine := scanner.Text()
	colors := strings.Fields(scanedLine)
	return colors
}
