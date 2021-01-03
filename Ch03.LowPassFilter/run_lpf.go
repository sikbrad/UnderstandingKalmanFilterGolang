package main

import (
	"fmt"
	"github.com/drgrib/iter"
	"github.com/sikbrad/UnderstandingKalmanFilterGolang/gqmathutil"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math/rand"
)

func GetSonar() float64 {

	stddev := 0.05
	w := 2.0 + stddev*rand.NormFloat64()

	return w
}

func LPF() func(x float64) float64 {
	prevX := 0.0
	alpha := 0.7
	isFirstRun := true

	return func(x float64) float64 {

		if isFirstRun {
			isFirstRun = false
			prevX = x //to remove bias
		}

		xlpf := alpha * prevX + (1 - alpha) * x
		prevX = xlpf

		return xlpf
	}
}

func main() {
	fmt.Println("Started LPF program")

	filter := LPF()
	nSamples := 100
	xSaved := gqmathutil.NewVectorZero(nSamples)  //avg x
	xmSaved := gqmathutil.NewVectorZero(nSamples) // measured x

	for k := range iter.N(nSamples) {
		xm := GetSonar()
		x := filter(xm)

		xSaved.SetVec(k, x)
		xmSaved.SetVec(k, xm)
	}

	dt := 0.02
	t := gqmathutil.Linspace(0, float64(nSamples)*dt, dt)

	gqmathutil.PrintMatrix(t, "t")
	gqmathutil.PrintMatrix(xSaved, "xSaved")
	gqmathutil.PrintMatrix(xmSaved, "xmSaved")

	p := gqmathutil.New2dPlotter("moving average filter(batch)")

	ptsX := gqmathutil.GetXyPointsFromVector(t, xSaved)
	ptsXm := gqmathutil.GetXyPointsFromVector(t, xmSaved)

	err := plotutil.AddLinePoints(p,
		"ptsXm", ptsXm,
		"ptsX", ptsX)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "tmp/res_images/ch03_lpf.png"); err != nil {
		panic(err)
	}
}
