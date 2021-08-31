package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func LShiftSource(a int, b int) int {
	if (a-1)%b == 0 {
		return a - b
	}
	return a
}

func UShiftSource(a int, b int, c int) int {
	if (a + b) > c {
		return a - b*(b-1)
	}
	return a + b
}
func writeMatrix(f *os.File, mtx [][]int, name string) {
	fmt.Fprintln(f, "\n"+name+": \n")
	for _, value := range mtx {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
}
func WriteToFile(filename string, shifftingStep int, a [][]int, b [][]int, c [][]int) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString("\n Shifting step: " + strconv.Itoa(shifftingStep) + "\n")
	writeMatrix(f, a, "A")
	writeMatrix(f, b, "B")
	writeMatrix(f, c, "C")

}

func WriteMatrix(filename string, marker string, c [][]int, d [][]int) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "\n"+marker+": \n")
	writeMatrix(f, c, "")
	writeMatrix(f, d, "")
}

type CBlockStruct struct {
	Num    int
	Matrix [][]int
}
