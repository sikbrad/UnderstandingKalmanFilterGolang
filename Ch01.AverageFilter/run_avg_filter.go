package main

import (
	"fmt"
	"github.com/sikbrad/UnderstandingKalmanFilterGolang/MathUtils"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"math/rand"
)

func AvgFilter() func(x float64) float64{
	prevAvg := float64(0)
	k := float64(1)

	return func(x float64) float64{
		alpha := float64((k - 1.0) / k)
		avg := alpha * prevAvg + (1.0 - alpha) * x

		prevAvg = avg
		k = k + 1

		return 0
	}
}

func GetVolt() float64{
	stddev := 4.0
	w := 0.0 + stddev * rand.NormFloat64()

	return w
}



func main()  {
	fmt.Println("Started AvgFilter")

	filter := AvgFilter()
	filter(10)

	//dt := 0.2

	t := MathUtils.Linspace(0,10,0.2)
	MathUtils.PrintMatrix(t, "t")

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "../tmp/res_images/points.png"); err != nil {
		panic(err)
	}
}
