package sequential

import (
	"fmt"
	"time"

	"github.com/MilicaPoparic/ntp/go/util"
)

func Sequential1(a [][]int, b [][]int, size int) {
	util.WriteMatrix("sequential.txt", "Matrices A, B", a, b)
	startTime := time.Now()
	var c [][]int
	// var c2 [][]int
	for i := 0; i < size; i++ {
		zeros := make([]int, size)
		c = append(c, zeros)
	}
	a, b = util.StepOne(a, b, size)

	c = util.AddAndMultiply(a, b, c, size)
	util.WriteToFile("sequential.txt", 1, a, b, c)

	for i := 1; i < size; i++ {
		for j := 0; j < size; j++ {
			a[j] = append(a[j][1:], a[j][:1]...)
		}
		b = append(b[1:], b[:1]...)
		c = util.AddAndMultiply(a, b, c, size)
		util.WriteToFile("sequential.txt", i+1, a, b, c)
	}

	elapsed := time.Since(startTime)
	for _, row := range c {
		fmt.Println(row)
	}
	fmt.Print("Process finished in ", elapsed)
}
