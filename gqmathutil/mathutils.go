package gqmathutil

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot/plotter"
)

func Linspace(min, max, spacing float64) *mat.VecDense {
	var arr []float64
	a := min
	for a < max {
		arr = append(arr, a)
		a = a + spacing
	}
	//arr = append(arr, max) //wtf?

	return mat.NewVecDense(len(arr), arr)
}

func PrintMatrix(mtx mat.Matrix, name string) {
	PrintMatrixWithFormatter(mtx, name, "%4.2f")
}

func PrintMatrixWithFormatter(mtx mat.Matrix, name, numFormat string) {
	rows, cols := mtx.Dims()
	fmt.Printf("prints matrix [%s], dim[%v,%v]\n", name, rows, cols)
	fa := mat.Formatted(mtx, mat.Prefix("    "), mat.Squeeze())
	fmtStr := fmt.Sprintf("%%s = %s\n", numFormat)
	//fmtStr := fmt.Sprintf("%s", numFormat)
	fmt.Printf(fmtStr, name, fa)
}

func ToFloatSlice(v *mat.VecDense) []float64 {
	retArr := make([]float64, (*v).Len())
	for i := range retArr {
		retArr[i] = v.AtVec(i)
	}
	return retArr
}

func GetXyPointsFromVector(xVec, yVec *mat.VecDense) plotter.XYs {
	return GetXyPointsFromFloatArray(
		ToFloatSlice(xVec),
		ToFloatSlice(yVec))
}

func GetXyPointsFromFloatArray(xArr, yArr []float64) plotter.XYs {
	if len(xArr) != len(yArr) {
		msg := fmt.Sprintf("xArr and yArr"+
			"length differes X[%v] Y[%v]",
			len(xArr), len(yArr))
		panic(msg)
	}
	n := len(xArr)
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = xArr[i]
		pts[i].Y = yArr[i]
		//if i == 0 {
		//	pts[i].X = rand.Float64()
		//} else {
		//	pts[i].X = pts[i-1].X + rand.Float64()
		//}
		//pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
