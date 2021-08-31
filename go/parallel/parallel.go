package parallel

import (
	"fmt"
	"math"
	"time"

	"github.com/MilicaPoparic/ntp/go/util"
)

func Parallel(a [][]int, b [][]int, size int, p int) {
	pSqrt := int(math.Sqrt(float64(p)))
	blockDim := size / pSqrt
	var dim1, step int
	dim2 := blockDim
	util.WriteMatrix("parallel.txt", "Matrices A, B: ", a, b)

	startTime := time.Now()

	allChans := make([][]chan []int, p)
	for i := range allChans {
		allChans[i] = make([]chan []int, 2) // za red i kolonu
		for j := range allChans[i] {
			allChans[i][j] = make(chan []int, 1)
		}
	}

	a, b = util.StepOne(a, b, size)
	var dest int
	cBlockHolder := make(chan util.CBlockStruct)

	for i := 0; i < size; i++ {
		var aBlock, bBlock [][]int
		var data [][][]int
		for j := dim1; j < dim2; j++ {
			aBlock = append(aBlock, a[j][step:step+blockDim])
			bBlock = append(bBlock, b[j][step:step+blockDim])
		}
		if len(aBlock[blockDim-1]) == blockDim {
			dest += 1
			if dest == p+1 {
				dest = 1
			}
			var cBlock [][]int
			for i := 0; i < blockDim; i++ {
				zeros := make([]int, blockDim)
				cBlock = append(cBlock, zeros)
			}
			data = append(data, aBlock)
			data = append(data, bBlock)
			data = append(data, cBlock)
			leftShiftSource := util.LShiftSource(dest+1, pSqrt)
			upShiftSource := util.UShiftSource(dest, pSqrt, p)
			sendChans := make([]chan []int, p)
			sendChans[0] = allChans[dest-1][0]            //dest levi shift
			sendChans[1] = allChans[dest-1][1]            // dest up shift
			sendChans[2] = allChans[leftShiftSource-1][0] // src levi shift
			sendChans[3] = allChans[upShiftSource-1][1]   // src up shift
			go RoutineJob(data, sendChans, size, blockDim, cBlockHolder, dest-1)
		}
		step += blockDim
		if (i+1)%blockDim == 0 {
			step = 0
			dim1 += blockDim
			dim2 += blockDim
		}
	}

	c := make([][]int, size)
	for i := range c {
		c[i] = make([]int, size)
		for j := range c[i] {
			c[i][j] = 0
		}
	}

	for t := 0; t < p; t++ {
		rp := <-cBlockHolder
		for k := 0; k < blockDim; k++ {
			i := rp.Num/pSqrt*blockDim + k
			j := rp.Num % pSqrt * blockDim
			c[i] = append(append(c[i][:j], rp.Matrix[k]...), c[i][j+blockDim:]...)
		}
	}
	for _, row := range c {
		fmt.Println(row)
	}

	elapsed := time.Since(startTime)
	fmt.Println("Process finished in: ", elapsed)
	util.WriteMatrix("parallel.txt", "Result: ", c, nil)
}

func RoutineJob(data [][][]int, chans []chan []int, size int, blockDim int, cBlockHolder chan util.CBlockStruct, source int) {
	for t := 0; t < size; t++ {
		util.AddAndMultiply(data[0], data[1], data[2], blockDim)
		if t == size-1 {
			util.WriteToFile("parallel.txt", source, data[0], data[1], data[2])
			cBlockHolder <- util.CBlockStruct{source, data[2]}
			break
		}
		d1 := make([]int, blockDim)
		for i := 0; i < blockDim; i++ {
			d1[i] = data[0][i][0]
		}

		chans[0] <- d1
		chans[1] <- data[1][0]
		s1 := <-chans[2]
		s2 := <-chans[3]

		for i := 0; i < blockDim; i++ {
			data[0][i] = append(data[0][i][1:], s1[i])
		}
		data[1] = append(data[1][1:], s2)

	}

}
