package conway

func fieldCopy(field [][]bool) [][]bool {
	cp := make([][]bool, len(field))
	for i := range cp {
		cp[i] = make([]bool, len(field[i]))
	}

	for i := range field {
		copy(cp[i], field[i])
	}

	return cp
}

func countNeighbours(field [][]bool, x, y int) int {
	count := 0

	if y+1 >= len(field) || y-1 < 0 {
		return -1
	}
	if x+1 >= len(field[0]) || x-1 < 0 {
		return -1
	}

	if field[y+1][x] {
		count++
	}
	if field[y+1][x+1] {
		count++
	}
	if field[y+1][x-1] {
		count++
	}
	if field[y][x+1] {
		count++
	}
	if field[y][x-1] {
		count++
	}
	if field[y-1][x] {
		count++
	}
	if field[y-1][x+1] {
		count++
	}
	if field[y-1][x-1] {
		count++
	}

	return count
}

func LifeGeneration(field [][]bool) {
	cp := fieldCopy(field)

	/*
	   Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	   Any live cell with two or three live neighbours lives on to the next generation.
	   Any live cell with more than three live neighbours dies, as if by overpopulation.
	   Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	*/

	for y := range cp {
		for x := range cp[y] {
			neighbours := countNeighbours(cp, x, y)

			if neighbours == -1 {
				continue
			}

			if cp[y][x] {
				if neighbours < 2 || neighbours > 3 {
					field[y][x] = false
				}
			} else {
				if neighbours == 3 {
					field[y][x] = true
				}
			}
		}
	}
}
