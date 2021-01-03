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

func SimpleKalman() func(x float64) float64 {
	A := 1.0
	H := 1.0
	Q := 0.0
	R := 4.0

	x := 0.0
	P := 0.0

	isFirstRun := true

	return func(z float64) float64 {

		if isFirstRun {
			isFirstRun = false
			x = 14
			P = 6
		}

		xp := A * x
		Pp := A * P / A + Q
		K := Pp * H / (H * Pp * H + R)
		x = xp + K * (z - H * xp)
		P = Pp - K * H * Pp

		return x
	}
}

func main() {
	fmt.Println("Started SimpleKalman program")

	filter := SimpleKalman()
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

	p := gqmathutil.New2dPlotter("simple kalman filter(1D)")

	ptsX := gqmathutil.GetXyPointsFromVector(t, xSaved)
	ptsXm := gqmathutil.GetXyPointsFromVector(t, xmSaved)

	err := plotutil.AddLinePoints(p,
		"ptsXm", ptsXm,
		"ptsX", ptsX)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "tmp/res_images/ch10_simple_kalman.png"); err != nil {
		panic(err)
	}
}

