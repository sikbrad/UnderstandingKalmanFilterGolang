package main

import (
	"fmt"
	"github.com/drgrib/iter"
	"github.com/sikbrad/UnderstandingKalmanFilterGolang/gqmathutil"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

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


func SimpleKalman2() func(x float64) (float64, float64, float64) {
	A := 1.0
	H := 1.0
	Q := 0.0
	R := 4.0

	x := 0.0
	P := 0.0

	isFirstRun := true

	return func(z float64) (float64, float64, float64) {

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

		return x, P, K
	}
}

func prac1() {
	fmt.Println("Started SimpleKalman1 program")

	filter := SimpleKalman()
	nSamples := 500
	xSaved := gqmathutil.NewVectorZero(nSamples) //avg x
	zSaved := gqmathutil.NewVectorZero(nSamples) // measured x

	for k := range iter.N(nSamples) {
		xm := gqmathutil.GetVolt()
		x := filter(xm)

		xSaved.SetVec(k, x)
		zSaved.SetVec(k, xm)
	}

	dt := 0.02
	t := gqmathutil.Linspace(0, float64(nSamples)*dt-dt, dt)

	gqmathutil.PrintMatrix(t, "t")
	gqmathutil.PrintMatrix(xSaved, "kf")
	gqmathutil.PrintMatrix(zSaved, "measurements")

	p := gqmathutil.New2dPlotter("simple kalman filter(1D)")

	ptsX := gqmathutil.GetXyPointsFromVectorDense(t, xSaved)
	ptsXm := gqmathutil.GetXyPointsFromVectorDense(t, zSaved)

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


// wanted to plot P and K, and show that it is converging
// but golang support for matrix .... really sucks.
func prac2() {
	//fmt.Println("Started SimpleKalman2 program")
	//
	//filter := SimpleKalman2()
	//nSamples := 500
	//xSaved := gqmathutil.NewMatrixZero(nSamples,3) //avg x
	//zSaved := gqmathutil.NewVectorZero(nSamples) // measured x
	//
	//for k := range iter.N(nSamples) {
	//	xm := gqmathutil.GetVolt()
	//	x, P, K := filter(xm)
	//
	//	xSaved.SetRow(k, []float64{x,P,K})
	//	zSaved.SetVec(k, xm)
	//}
	//
	//dt := 0.02
	//t := gqmathutil.Linspace(0, float64(nSamples)*dt-dt, dt)


	//p := gqmathutil.New2dPlotter("simple kalman filter(1D) P")
	//
	//tt := mat.NewDense(t.Len(),1, gqmathutil.ToFloatSlice(t))
	//ptsX := gqmathutil.GetXyPointsFromDense(tt, xSaved.ColView(1).RawVector())
	//ptsXm := gqmathutil.GetXyPointsFromVectorDense(t, zSaved)
	//
	//err := plotutil.AddLinePoints(p,
	//	"ptsXm", ptsXm,
	//	"ptsX", ptsX)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Save the plot to a PNG file.
	//if err := p.Save(4*vg.Inch, 4*vg.Inch, "tmp/res_images/ch10_simple_kalman_p.png"); err != nil {
	//	panic(err)
	//}
}


func main(){
	prac1()
	prac2()
}

