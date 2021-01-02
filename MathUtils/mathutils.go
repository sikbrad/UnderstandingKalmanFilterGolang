package MathUtils

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
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

func PrintMatrix(mtx mat.Matrix, name string){
	PrintMatrixWithFormatter(mtx,name,"%4.2f")
}

func PrintMatrixWithFormatter(mtx mat.Matrix, name, numFormat string){
	fa := mat.Formatted(mtx, mat.Prefix("    "), mat.Squeeze())
	fmtStr := fmt.Sprintf("%%s = %s\n", numFormat)
	fmt.Printf(fmtStr, name, fa)
}