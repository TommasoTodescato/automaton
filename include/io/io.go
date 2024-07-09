package io

import (
	"fmt"
)

// warning : this only works in unix-like for now
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintField(field [][]bool) {

	clearScreen()

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] {
				fmt.Print("■ ")
			} else {
				fmt.Print("· ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
