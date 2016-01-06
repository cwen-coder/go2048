package go2048

import (
	//	"fmt"
	"github.com/nsf/termbox-go"
)

var Score int
var step int

type Status uint

func coverPrintStr(x, y int, str string, fg, bg termbox.Attribute) error {
	xx := x
	for n, c := range str {
		if c == '\n' {
			y++
			xx = x - n - 1
		}
		termbox.SetCell(xx+n, y, c, fg, bg)
	}
	termbox.Flush()
	return nil
}

const (
	Win Status = iota
	Lose
	Add
	Max = 2048
)
