package sequential

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/MilicaPoparic/ntp/go/util"
	"github.com/barkimedes/go-deepcopy"
)

func saveBlocks(blocks [][][][]int, data [][][]int, index int, blockDim int) [][][][]int {
	for i := 0; i < 3; i++ {
		for k := 0; k < blockDim; k++ {
			blocks[index][i][k] = data[i][k]
		}
	}
	return blocks
}

func Sequential(a [][]int, b [][]int, n int, p int) {
	pSqrt := int(math.Sqrt(float64(p)))
	blockDim := n / pSqrt
	blocks := make([][][][]int, p)
	for i := range blocks {
		blocks[i] = make([][][]int, 3)
		for j := range blocks[i] {
			blocks[i][j] = make([][]int, blockDim)
			for k := range blocks[i][j] {
				blocks[i][j][k] = make([]int, blockDim)
			}
		}
	}

	cBlocks := make([][][]int, p)
	for i := range cBlocks {
		cBlocks[i] = make([][]int, blockDim)
		for j := range cBlocks[i] {
			cBlocks[i][j] = make([]int, blockDim)
		}
	}
	// util.WriteMatrix("sequential.txt", "Matrices A, B", a, b)
	startTime := time.Now()

	cBlock := make([][]int, blockDim)
	for i := 0; i < blockDim; i++ {
		zeros := make([]int, blockDim)
		cBlock[i] = append(zeros)
	}

	dest := 0
	a, b = util.StepOne(a, b, n)
	var aBlock, bBlock [][]int
	var data [][][]int
	for i := 0; i < n; i += blockDim {
		mtxA := a[i : i+blockDim]
		mtxB := b[i : i+blockDim]
		for j := 0; j < n; j += blockDim {
			for k := range mtxA {
				aBlock = append(aBlock, mtxA[k][j:j+blockDim])
				bBlock = append(bBlock, mtxB[k][j:j+blockDim])
				if len(aBlock) == blockDim {
					data = append(data, aBlock)
					data = append(data, bBlock)
					data = append(data, cBlock)
					blocks = saveBlocks(blocks, data, dest, blockDim)
					dest += 1
					bBlock = bBlock[:0]
					aBlock = aBlock[:0]
					data = data[:0]
				}
			}
		}
	}

	type Foo struct {
		B [][][][]int
	}

	x := Foo{
		B: blocks,
	}

	// endSeq := time.Since(startTime)
	for m := 0; m < n; m++ {
		bs := deepcopy.MustAnything(x)
		blocksShifted := bs.(Foo).B

		sftL := deepcopy.MustAnything(x)
		blockPerShiftingL := sftL.(Foo).B

		sftU := deepcopy.MustAnything(x)
		blockPerShiftingU := sftU.(Foo).B

		for r := 0; r < p; r++ {
			process := r + 1
			util.AddAndMultiply1(x.B[r][0], x.B[r][1], x.B[r][2], blockDim)
			// util.WriteToFile("sequential.txt", m+process, x.B[r][0], x.B[r][1], x.B[r][2])
			for s := 0; s < blockDim; s++ {
				for k := 0; k < blockDim; k++ {
					cBlocks[r][s][k] += x.B[r][2][s][k]
				}
			}
			leftShiftDest := util.LShiftDest(process, pSqrt)

			newCol := make([]int, blockDim)
			for p := 0; p < blockDim; p++ {
				newCol[p] = x.B[leftShiftDest-1][0][p][0]
			}

			for j := 0; j < blockDim; j++ {
				blocksShifted[r][0][j] = append(blockPerShiftingL[r][0][j][1:], newCol[j])
			}

			upShiftDest := util.UShiftDest(process, pSqrt)
			newRow := (x.B[upShiftDest-1][1][0])
			blocksShifted[r][1] = append(blockPerShiftingU[r][1][1:], newRow)
		}
		x.B = blocksShifted
	}

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = 0
		}
	}

	for t := 0; t < p; t++ {
		for k := 0; k < blockDim; k++ {
			i := t/pSqrt*blockDim + k
			j := t % pSqrt * blockDim
			result[i] = append(append(result[i][:j], cBlocks[t][k]...), result[i][j+blockDim:]...)
		}
	}
	elapsed := time.Since(startTime)
	for _, row := range result {
		fmt.Println(row)
	}
	fmt.Print("Process finished in ", elapsed)
	f, err := os.OpenFile("resources/sequentialWeek25.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	f.WriteString(fmt.Sprint(elapsed.Seconds()) + "\n")
	defer f.Close()

	// paralelTime := elapsed.Seconds() - endSeq.Seconds()
	// fmt.Println("sequential time", endSeq.Seconds())
	// fmt.Println("parallel time", paralelTime/elapsed.Seconds())
}
