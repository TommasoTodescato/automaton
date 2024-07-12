package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func InitScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return screen
}

func PrintField(screen tcell.Screen, field [][]bool) {
	screen.Clear()

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			c := '·'
			if field[y][x] {
				c = '■'
			}
			screen.SetContent(x, y, c, nil, tcell.StyleDefault)
		}
	}
	screen.Show()
}
