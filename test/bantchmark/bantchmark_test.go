package main

import (
	"testing"

	"github.com/iicecookie/YabbiTest/internal/models/colorTable"
)

func BenchmarkColorTableInit(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := colorTable.New(5, 5, "Red", "Green", "Blue")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColorTable(b *testing.B) {

	for i := 0; i < b.N; i++ {

		b.StopTimer()
		table, err := colorTable.New(5, 5, "Red", "Green", "Blue")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		table.GetMaxNeighboringColors()
	}
}
