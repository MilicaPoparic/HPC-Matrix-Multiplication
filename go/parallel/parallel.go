package parallel

import (
	"fmt"
	"math"
	"time"

	"github.com/MilicaPoparic/ntp/go/util"
)

func routine(data [][][]int, chans []chan []int, size int, blockDim int) {
	//treba da saljem na neki kanal ove isparcane rezultate i da sklopim matricu posle
	for t := 0; t < size; t++ {
		util.AddAndMultiply(data[0], data[1], data[2], blockDim)
		if t == size-1 {
			break
		}
		d1 := make([]int, blockDim)
		for i := 0; i < blockDim; i++ {
			d1[i] = data[0][i][0]
		}

		chans[0] <- d1
		chans[1] <- data[1][0]
		s1 := <-chans[2]
		fmt.Println(d1, "DESTINACIJA ", s1, " SOURCE ")
		for i := 0; i < blockDim; i++ {
			data[0][i] = append(data[0][i][1:], s1[i])

		}
		s2 := <-chans[3]
		data[1] = append(data[1][1:], s2)

	}
}

func Parallel(a [][]int, b [][]int, size int, p int) {
	pSqrt := int(math.Sqrt(float64(p)))
	blockDim := size / pSqrt
	var dim1, step int
	dim2 := blockDim
	var cBlock [][]int
	for i := 0; i < blockDim; i++ {
		zeros := make([]int, blockDim)
		cBlock = append(cBlock, zeros)
	}

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
	//treba mi lista sa 4 kanala,
	//prva dva su dva source, druga dva 2 dest npr
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
			data = append(data, aBlock)
			data = append(data, bBlock)
			data = append(data, cBlock)
			leftShiftSource := util.LShiftSource(dest+1, pSqrt)
			upShiftSource := util.UShiftSource(dest, pSqrt, p)
			var sendChans []chan []int
			sendChans = append(sendChans, allChans[dest-1][0])            //dest levi shift
			sendChans = append(sendChans, allChans[dest-1][1])            // dest up shift
			sendChans = append(sendChans, allChans[leftShiftSource-1][0]) // src levi shift
			sendChans = append(sendChans, allChans[upShiftSource-1][1])   // src up shift
			go routine(data, sendChans, size, blockDim)
		}
		step += blockDim
		if (i+1)%blockDim == 0 {
			step = 0
			dim1 += blockDim
			dim2 += blockDim
		}
	}
	elapsed := time.Since(startTime)
	fmt.Println("Process finished in: ", elapsed)
}
