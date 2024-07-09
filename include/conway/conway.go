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

func GameOfLifeIteration(field [][]bool) {
	cp := fieldCopy(field)

	for y := range cp {
		for x := range cp[y] {
			if cp[y][x] {
				field[y][x+1] = true
				field[y][x] = false
			}
		}
	}
}
