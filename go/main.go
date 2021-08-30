package main

import (
	"github.com/MilicaPoparic/ntp/go/parallel"
)

func main() {

	a := [][]int{
		{15, -11, -12, 12},
		{-15, -2, 15, -15},
		{12, 14, -12, -6},
		{-1, -8, 16, -13},
	}
	b := [][]int{
		{0, 15, 14, 9},
		{-3, -7, -12, -4},
		{10, 10, -16, 15},
		{-13, -3, 9, 3},
	}
	parallel.Parallel(a, b, 4, 4)
	// sequential.Sequential(a, b, 4)
}
