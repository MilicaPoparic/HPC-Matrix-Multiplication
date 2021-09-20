package parallel

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/MilicaPoparic/ntp/go/util"
)

func Parallel(a [][]int, b [][]int, size int, p int) {
	pSqrt := int(math.Sqrt(float64(p)))
	blockDim := size / pSqrt
	// util.WriteMatrix("parallel.txt", "Matrices A, B: ", a, b)

	startTime := time.Now()

	allChans := make([][]chan []int, p)
	for i := range allChans {
		allChans[i] = make([]chan []int, 2) // row, col
		for j := range allChans[i] {
			allChans[i][j] = make(chan []int, blockDim) //row, col dim blockDim
		}
	}

	a, b = util.StepOne(a, b, size)
	dest := 0
	cBlockHolder := make(chan util.CBlockStruct)

	aBlock := make([][]int, blockDim)
	bBlock := make([][]int, blockDim)

	for i := 0; i < size; i += blockDim {
		mtxA := a[i : i+blockDim]
		mtxB := b[i : i+blockDim]
		for j := 0; j < size; j += blockDim {
			for k := range mtxA {
				aBlock[k] = mtxA[k][j : j+blockDim]
				bBlock[k] = mtxB[k][j : j+blockDim]
				if len(aBlock[blockDim-1]) == blockDim {
					// intialize c blok with all zeros
					cBlock := make([][]int, blockDim)
					for i := range cBlock {
						cBlock[i] = make([]int, blockDim)
					}

					data := make([][][]int, 3)
					data[0] = aBlock
					data[1] = bBlock
					data[2] = cBlock

					dest += 1

					aBlock = make([][]int, blockDim)
					bBlock = make([][]int, blockDim)

					leftShiftSource := util.LShiftSource(dest+1, pSqrt)
					upShiftSource := util.UShiftSource(dest, pSqrt, p)

					sendChans := make([]chan []int, 4)
					sendChans[0] = allChans[dest-1][0]            //dest left shift
					sendChans[1] = allChans[dest-1][1]            // dest up shift
					sendChans[2] = allChans[leftShiftSource-1][0] // src left shift
					sendChans[3] = allChans[upShiftSource-1][1]   // src up shift
					go RoutineJob(data, sendChans, size, blockDim, cBlockHolder, dest-1)
				}
			}
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
	f, err := os.OpenFile("resources/parallelWeak25.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	f.WriteString(fmt.Sprint(elapsed.Seconds()) + "\n")
	defer f.Close()
}

func RoutineJob(data [][][]int, chans []chan []int, size int, blockDim int, cBlockHolder chan util.CBlockStruct, source int) {
	for t := 0; t < size; t++ {
		util.AddAndMultiply(data[0], data[1], data[2], blockDim)
		if t == size-1 {
			// util.WriteToFile("parallel.txt", source, data[0], data[1], data[2])
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
