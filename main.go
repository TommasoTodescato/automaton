package main

import (
	conway "automaton/include/conway"
	io "automaton/include/io"
	"time"
)

func main() {
	dy, dx := 10, 10

	f := make([][]bool, dy)
	for i := range f {
		f[i] = make([]bool, dx)
	}

	f[0][0] = true
	f[2][0] = true

	for i := 0; i < dx-1; i++ {
		io.PrintField(f)

		conway.GameOfLifeIteration(f)
		time.Sleep(1 * time.Second)

		io.PrintField(f)
	}
}
