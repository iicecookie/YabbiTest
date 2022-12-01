package colorTable

import (
	"github.com/TwiN/go-color"
)

type Color int

const (
	Gray Color = iota
	Black
	Red
	Green
	Yellow
	Blue
)

var (
	ColorMap = map[string]Color{
		"Black":  Black,
		"Red":    Red,
		"Green":  Green,
		"Yellow": Yellow,
		"Blue":   Blue,
		"Gray":   Gray,
	}
)

func (c Color) GetCode() string {
	return [...]string{color.Gray, color.Black, color.Red, color.Green, color.Yellow, color.Blue}[c]
}
