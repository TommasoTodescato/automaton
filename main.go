package main

import (
	conway "automaton/include/conway"
	io "automaton/include/io"
	"time"
)

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

	for i := 0; i < 100; i++ {
		io.PrintField(f)

		conway.LifeGeneration(f)
		time.Sleep(500 * time.Millisecond)

		io.PrintField(f)
	}
}
