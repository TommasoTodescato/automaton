package main

import (
	"automaton/include/conway"
	"automaton/include/field"
	"automaton/include/io"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	f := field.InitField(30, 30)

	screen := io.InitScreen()
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
					if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
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
			io.PrintField(screen, f)
		case <-quit:
			return
		}
	}
}
