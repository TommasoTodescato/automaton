package io

import (
	"fmt"
)

func print_field(field [][]bool) {

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {

			if field[y][x] {
				fmt.Print("X ")
			} else {
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}

}

func Test() {

	dx, dy := 5, 5

	a := make([][]bool, dy)

	for i := range a {
		a[i] = make([]bool, dx)
	}

	a[0][2] = true
	a[0][1] = true

	print_field(a)
}
