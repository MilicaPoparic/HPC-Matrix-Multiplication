package sequential

import (
	"fmt"
	"time"

	"github.com/MilicaPoparic/ntp/go/util"
)

func Sequential(a [][]int, b [][]int, size int) {
	startTime := time.Now()
	var c [][]int
	for i := 0; i < size; i++ {
		zeros := make([]int, size)
		c = append(c, zeros)
	}
	a, b = util.StepOne(a, b, size)
	c = util.AddAndMultiply(a, b, c, size)
	fmt.Println(c)
	for i := 1; i < size; i++ {
		for j := 0; j < size; j++ {
			a[j] = append(a[j][1:], a[j][:1]...)
		}
		b = append(b[1:], b[:1]...)
		c = util.AddAndMultiply(a, b, c, size)
	}

	elapsed := time.Since(startTime)
	for _, row := range c {
		fmt.Println(row)
	}
	fmt.Print("Process finished in ", elapsed)
}
