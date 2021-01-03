package gqmathutil

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"math/rand"
	"os"
	"strconv"
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
		msg := fmt.Sprintf(
			"xArr and yArr"+
				"length differes X[%v] Y[%v]",
			len(xArr), len(yArr))
		panic(msg)
	}
	n := len(xArr)
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = xArr[i]
		pts[i].Y = yArr[i]
	}
	return pts
}

func New2dPlotter(name string) *plot.Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	return p
}

func NewVectorOne(n int) *mat.VecDense {
	return newVectorWithSingleValue(n, 1.0)
}
func NewVectorZero(n int) *mat.VecDense {
	return newVectorWithSingleValue(n, 0.0)
}

func newVectorWithSingleValue(n int, val float64) *mat.VecDense {
	vec := mat.NewVecDense(n, nil)

	for i := range make([]int, vec.Len()) {
		vec.SetVec(i, val)
	}

	return vec
}

func SumVector(vec *mat.VecDense) float64 {
	sum := 0.0
	for i := range make([]int, vec.Len()) {
		sum += vec.AtVec(i)
	}
	return sum
}


func SonarDataLoader() (func() float64, error){
	csvFile, err := os.Open("data/sonarAlt.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	idxData := 0

	return func() float64{
		elem := csvData[0][idxData]
		idxData++
		datF, _ := strconv.ParseFloat(elem, 64)
		return datF
	}, nil
}

func GetVolt() float64 {
	stddev := 4.0
	w := 0.0 + stddev*rand.NormFloat64()

	return w
}