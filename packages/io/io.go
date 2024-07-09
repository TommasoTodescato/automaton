package io

import (
	"fmt"
)

func PrintField(field [][]bool) {

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
	fmt.Println()
}
