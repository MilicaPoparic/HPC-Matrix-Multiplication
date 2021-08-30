package util

func AddAndMultiply(a [][]int, b [][]int, c [][]int, size int) [][]int {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c[i][j] += a[i][j] * b[i][j]
		}
	}
	return c
}

func StepOne(a [][]int, b [][]int, size int) ([][]int, [][]int) {
	for i := 0; i < size; i++ {
		a[i] = append(a[i][i:], a[i][:i]...)
		var col []int
		var shifted []int
		for j := 0; j < size; j++ {
			col = append(col, b[j][i])
		}
		shifted = append(col[i:], col[:i]...)
		for j := 0; j < size; j++ {
			b[j][i] = shifted[j]
		}
	}
	return a, b
}

func LShiftDest(a int, b int) int {
	if a%b == 0 {
		return a + b
	}
	return a
}

func LShiftSource(a int, b int) int {
	if (a-1)%b == 0 {
		return a - b
	}
	return a
}
func UShiftDest(a int, b int) int {
	if (a - b) > 0 {
		return a - b
	}
	return a + b*(b-1)
}

func UShiftSource(a int, b int, c int) int {
	if (a + b) > c {
		return a - b*(b-1)
	}
	return a + b
}
