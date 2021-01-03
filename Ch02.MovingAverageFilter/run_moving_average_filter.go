package main

import (
	"fmt"
	"github.com/drgrib/iter"
	"github.com/sikbrad/UnderstandingKalmanFilterGolang/gqmathutil"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func MovAvgFilter(windowSize int) func(x float64) float64 {
	prevAvg := 0.0
	k := 1.0
	n := windowSize
	xbuf := gqmathutil.NewVectorOne(n + 1)
	isFirstRun := true

	return func(x float64) float64 {

		if isFirstRun {
			isFirstRun = false
			for idx, _ := range gqmathutil.ToFloatSlice(xbuf) {
				xbuf.SetVec(idx, x)
			}
			prevAvg = x //to rem bias
		}

		for m := range iter.N(n) {
			xbuf.SetVec(m, xbuf.AtVec(m+1))
		}
		xbuf.SetVec(n, x)

		avg := prevAvg + (x-xbuf.AtVec(0))/float64(n)

		prevAvg = avg
		k = k + 1

		return avg
	}
}

func main() {
	fmt.Println("Started MovingAverageFilter program")

	filter := MovAvgFilter(50)
	//nSamples := 1501
	nSamples := 500
	xSaved := gqmathutil.NewVectorZero(nSamples)  //avg x
	xmSaved := gqmathutil.NewVectorZero(nSamples) // measured x

	dataLoader, err := gqmathutil.SonarDataLoader()
	if err!=nil{
		panic("cannot open sonar data file")
	}

	for k := range iter.N(nSamples) {
		xm := dataLoader()
		x := filter(xm)

		xSaved.SetVec(k, x)
		xmSaved.SetVec(k, xm)
	}

	dt := 0.02
	t := gqmathutil.Linspace(0, float64(nSamples)*dt-dt, dt)

	gqmathutil.PrintMatrix(t, "t")
	gqmathutil.PrintMatrix(xSaved, "xSaved")
	gqmathutil.PrintMatrix(xmSaved, "xmSaved")

	p := gqmathutil.New2dPlotter("moving average filter(batch)")

	ptsX := gqmathutil.GetXyPointsFromVectorDense(t, xSaved)
	ptsXm := gqmathutil.GetXyPointsFromVectorDense(t, xmSaved)

	err = plotutil.AddLinePoints(p,
		"ptsXm", ptsXm,
		"ptsX", ptsX)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "tmp/res_images/ch02_average_filter_points_recursive.png"); err != nil {
		panic(err)
	}
}
