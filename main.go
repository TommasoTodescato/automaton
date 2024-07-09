package main

import (
	io "automaton/packages/io"
)

func main() {
	dy, dx := 5, 5

	f := make([][]bool, dy)
	for i := range f {
		f[i] = make([]bool, dx)
	}

	io.PrintField(f)

	f[0][3] = true
	f[2][4] = true

	io.PrintField(f)
}
