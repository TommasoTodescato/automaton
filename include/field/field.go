package field

func InitField(dx, dy int) [][]bool {
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

	return f
}
