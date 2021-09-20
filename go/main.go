package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// a := [][]int{
	// 	{-14, 10, -11, 2, 7, -21},
	// 	{30, 5, 33, -34, 12, 35},
	// 	{25, 4, 21, 26, -26, -6},
	// 	{16, 36, -7, 0, 13, -11},
	// 	{-17, 17, 3, -3, -34, -6},
	// 	{2, -1, -10, 9, -14, 14},
	// }
	// b := [][]int{
	// 	{-9, 2, -16, 0, -32, -9},
	// 	{-21, -6, 10, -12, 10, 18},
	// 	{-9, 1, 9, 35, -21, 34},
	// 	{2, 28, -5, 30, -1, -29},
	// 	{-1, 20, -1, 27, -29, -1},
	// 	{-6, -36, -25, -16, -7, -33},
	// }
	// a := [][]int{
	// 	{15, -11, -12, 12},
	// 	{-15, -2, 15, -15},
	// 	{12, 14, -12, -6},
	// 	{-1, -8, 16, -13},
	// }
	// b := [][]int{
	// 	{0, 15, 14, 9},
	// 	{-3, -7, -12, -4},
	// 	{10, 10, -16, 15},
	// 	{-13, -3, 9, 3},
	// }
	// n := 500
	// for i := 0; i < 30; i++ {
	// 	a := make([][]int, n)
	// 	for i := range a {
	// 		a[i] = make([]int, n)
	// 		for j := range a[i] {
	// 			a[i][j] = rand.Intn(100-0) + 0

	// 		}
	// 	}
	// 	b := make([][]int, n)
	// 	for i := range b {
	// 		b[i] = make([]int, n)

	// 		for j := range b[i] {
	// 			b[i][j] = rand.Intn(100-0) + 0
	// 		}
	// 	}

	// 	sequential.Sequential(a, b, 500, 1)
	// 	// parallel.Parallel(a, b, 500, 25)
	// }
	cpuCount := []int{4, 16, 25}
	speedup := []float64{3.03, 2.03, 1.78}
	maxSpeedup := []float64{}

	//calculate amdal
	// for c := range cpuCount {
	// 	maxSpeedup = append(maxSpeedup, 1/(0.01+float64(0.99)/float64(cpuCount[c])))
	// }
	//calculate gustaf
	for c := range cpuCount {
		maxSpeedup = append(maxSpeedup, 0.01+float64(0.99)*float64(cpuCount[c]))
	}

	plotGraph(speedup, maxSpeedup, cpuCount, "Slabo")

}

func plotGraph(speedup []float64, maxSpeedup []float64, cpuCount []int, t string) {
	speedup_points := make(plotter.XYs, len(speedup))
	max_points := make(plotter.XYs, len(maxSpeedup))

	for i := range speedup {
		speedup_points[i].X = float64(cpuCount[i])
		speedup_points[i].Y = speedup[i]
		max_points[i].X = float64(cpuCount[i])
		max_points[i].Y = maxSpeedup[i]
	}

	p := plot.New()

	p.Title.Text = t + " skaliranje"
	p.X.Label.Text = "Broj procesora"
	p.Y.Label.Text = "Ubrzanje"

	err := plotutil.AddLinePoints(p,
		"", speedup_points,
		"", max_points,
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "resources/weak_scaling.png"); err != nil {
		panic(err)
	}

}
