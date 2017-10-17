package Approximation

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

func TestSimpleFunction(t *testing.T) {
	input := [][]float64{{-2}, {-1}, {0}, {1}, {2}}
	output := []float64{4, 1, 0, 1, 4}

	net := GetRBFNet(input, output)
	t.Log(net)

	t.Log(net.Predict([]float64{3}))
	t.Log(net.Predict([]float64{4}))
	t.Log(net.Predict([]float64{5}))

	file, _ := os.Create("quadraticTest.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	cx := -20.0
	writer.Write([]string{"X", "Y"})
	for cx <= 20 {

		line := []string{toString(cx), toString(net.Predict([]float64{cx}))}
		writer.Write(line)

		cx += 0.1
	}

}

func TestRBFApproximation(t *testing.T) {
	/*approximationSlice := []float64{
	2534.822532831742,
	2983.7148097888216,
	3199.186551601226,
	3772.97825043461,
	5613.178073327571,
	9004.880056043232,
	9343.201089385899,
	9552.019187813006,
	14617.453209491208,
	15390.952931039958,
	15666.20698139584}*/

	//approximationMatrix := mat64.NewDense(11, 1, approximationSlice)
	//value :=
	//testFitness = ApproximateRBF(approximationMatrix, value []float64) (fitness float64) {

}

func toString(num float64) string {
	return strconv.FormatFloat(num, 'f', 6, 64)
}
