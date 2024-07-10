package main

import (
	"automaton/include/conway"
	"automaton/include/io"
	"time"
)

func main() {
	dy, dx := 20, 20

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

	for i := 0; i < 100; i++ {
		time.Sleep(300 * time.Millisecond)
		io.PrintField(f)
		conway.LifeGeneration(f)
	}
}
