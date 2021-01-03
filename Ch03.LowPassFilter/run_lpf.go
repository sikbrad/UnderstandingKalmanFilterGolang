package main

import (
	"fmt"
	"github.com/drgrib/iter"
	"github.com/sikbrad/UnderstandingKalmanFilterGolang/gqmathutil"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)


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

	p := gqmathutil.New2dPlotter("LPF")

	ptsX := gqmathutil.GetXyPointsFromVectorDense(t, xSaved)
	ptsXm := gqmathutil.GetXyPointsFromVectorDense(t, xmSaved)

	err = plotutil.AddLinePoints(p,
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
