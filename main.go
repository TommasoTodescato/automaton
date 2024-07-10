package main

import (
	"automaton/include/conway"
	"automaton/include/field"
	"automaton/include/io"
	"time"
)

func main() {
	f := field.InitField(30, 30)

	screen := io.InitScreen()
	if screen == nil {
		return
	}
	defer screen.Fini()

	for i := 0; i < 70; i++ {
		io.PrintField(screen, f)
		conway.LifeGeneration(f)
		time.Sleep(200 * time.Millisecond)
	}
}
