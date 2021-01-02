package main

import (
	"fmt"
	"github.com/drgrib/iter"
	"github.com/sikbrad/UnderstandingKalmanFilterGolang/gqmathutil"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math/rand"
)

func AvgFilter() func(x float64) float64 {
	prevAvg := float64(0)
	k := float64(1)

	return func(x float64) float64 {
		alpha := float64((k - 1.0) / k)
		avg := alpha*prevAvg + (1.0-alpha)*x

		prevAvg = avg
		k = k + 1

		return avg
	}
}

func GetVolt() float64 {
	stddev := 4.0
	w := 0.0 + stddev*rand.NormFloat64()

	return w
}

func main() {
	fmt.Println("Started AvgFilter")

	avgFilter := AvgFilter()
	avgFilter(10)

	dt := 0.2
	t := gqmathutil.Linspace(0, 10, dt)
	gqmathutil.PrintMatrix(t, "t")

	nSamples := t.Len()
	avgSaved := mat.NewVecDense(nSamples, nil)
	xmSaved := mat.NewVecDense(nSamples, nil)

	for i := range iter.N(nSamples) {
		xm := GetVolt()
		avg := avgFilter(xm)

		avgSaved.SetVec(i, avg)
		xmSaved.SetVec(i, xm)
	}

	gqmathutil.PrintMatrix(avgSaved, "avgSaved")
	gqmathutil.PrintMatrix(xmSaved, "xmSaved")

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	//pts := gqmathutil.GetXyPointsFromFloatArray(
	//	gqmathutil.ToFloatSlice(t),
	//	gqmathutil.ToFloatSlice(avgSaved))

	ptsAvg := gqmathutil.GetXyPointsFromVector(t, avgSaved)
	ptsXm := gqmathutil.GetXyPointsFromVector(t, xmSaved)

	err = plotutil.AddLinePoints(p,
		"ptsAvg", ptsAvg,
		"ptsXm", ptsXm)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "../tmp/res_images/points.png"); err != nil {
		panic(err)
	}
}
