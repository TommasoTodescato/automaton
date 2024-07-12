package main

import (
	"automaton/engine/conway"
	"automaton/tui"
	"automaton/web"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

func run_tui(f [][]bool) {

	screen := tui.InitScreen()
	if screen == nil {
		return
	}
	defer screen.Fini()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	quit := make(chan struct{})
	go func() {
		for {
			if ev := screen.PollEvent(); ev != nil {
				switch ev := ev.(type) {
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
						close(quit)
						return
					}
				}
			}
		}
	}()

	for i := 0; i < 70; i++ {
		select {
		case <-ticker.C:
			conway.LifeGeneration(f)
			tui.PrintField(screen, f)
		case <-quit:
			return
		}
	}
}

func main() {

	dy, dx := 30, 30

	f := make([][]bool, dy)
	for i := range f {
		f[i] = make([]bool, dx)
	}
	f[2][3] = true
	f[3][4] = true
	f[4][2] = true
	f[4][3] = true
	f[4][4] = true
	f[5][5] = true

	mode := -1
	fmt.Print("Choose the UI: \n\n0-Console\n1-Browser\n<other>-Quit\n")
	fmt.Scan(&mode)
	if mode == 0 {
		run_tui(f)
	} else if mode == 1 {
		web.Run()
	} else {
		return
	}

}
