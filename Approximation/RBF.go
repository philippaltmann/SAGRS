package Approximation

import (
	"fmt"
	"math"

	"github.com/gonum/matrix/mat64"
)

type BiasFunction struct {
	center []float64
	width  float64
}

type RBFApproximator struct {
	HiddenLayer []BiasFunction
	Weights     []float64
}

/*func GenerateRBFApproximator(size, dimensions int) *RBFApproximator {
	var hiddenLayer []BiasFunction
	for i := 0; i < size; i++ {
		biasFunction := BiasFunction{make([]float64, dimensions), 0.0}
		hiddenLayer = append(hiddenLayer, biasFunction)
	}
	weights := make([]float64, size)
	return &RBFApproximator{hiddenLayer, weights}
}*/

func (f BiasFunction) Calculate(value []float64) float64 {
	//return math.Sqrt((1 + (f.width * EuclideanDistance(value, f.center))))
	//return math.Exp(-(EuclideanDistance(value, f.center)) / (2 * math.Pow(f.width, 2)))
	return 1 - math.Exp(-(EuclideanDistance(value, f.center))/(2*math.Pow(f.width, 2)))
}

func EuclideanDistance(from []float64, to []float64) (distance float64) {
	for i, v := range from {
		distance += math.Pow((v - to[i]), 2)
	}
	return
}

func (net RBFApproximator) GetFunctionMatrix(input [][]float64) *mat64.Dense {
	var slice []float64
	for _, layer := range net.HiddenLayer {
		for _, value := range input {
			slice = append(slice, layer.Calculate(value))
		}
	}
	return mat64.NewDense(len(net.HiddenLayer), len(input), slice)
}

func GetRBFApproximator(input [][]float64, output []float64) RBFApproximator {
	net := new(RBFApproximator)

	var averageDistance float64
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			averageDistance += EuclideanDistance(input[i], input[j])
		}
	}
	averageDistance = math.Sqrt(averageDistance)
	averageDistance /= float64(len(input))

	//"Train" Layers
	for i := 0; i < len(input); i++ {
		bf := new(BiasFunction)
		bf.center = input[i] //TODO clone
		bf.width = averageDistance
		net.HiddenLayer = append(net.HiddenLayer, *bf)
	}

	//"Train" Weights
	functionMatrix := net.GetFunctionMatrix(input)
	functionMatrix.Inverse(functionMatrix)

	var W mat64.Dense
	W.Mul(functionMatrix, mat64.NewDense(len(output), 1, output))

	net.Weights = mat64.Col(net.Weights, 0, &W)
	//fmt.Println(mat64.Formatted(&W))

	return *net
}
func (net *RBFApproximator) UpdateRBFApproximator(input [][]float64, output []float64) {
	/*if len(input) != len(net.HiddenLayer) {
		panic("Mismatching Input Size")
	}
	if len(input[0]) != len(net.HiddenLayer[0].center) {
		panic("Mismatching Dimensions")
	}*/
	var averageDistance float64
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			averageDistance += EuclideanDistance(input[i], input[j])
		}
	}
	averageDistance = math.Sqrt(averageDistance)
	averageDistance /= float64(len(input))
	//averageDistance *= float64(len(output))
	//averageDistance *= 4
	//averageDistance = 10.0

	//"Train" Layers
	net.HiddenLayer = make([]BiasFunction, len(input))
	for i := 0; i < len(input); i++ {
		net.HiddenLayer[i].center = append([]float64(nil), input[i]...)
		net.HiddenLayer[i].width = averageDistance
	}

	//"Train" Weights
	functionMatrix := net.GetFunctionMatrix(input)
	functionMatrix.Inverse(functionMatrix)

	var W mat64.Dense
	W.Mul(functionMatrix, mat64.NewDense(len(output), 1, output))

	net.Weights = mat64.Col(net.Weights, 0, &W)
	//fmt.Printf("Updated RBF Hidden Layer by %d Values\n", updated)
	//fmt.Println(mat64.Formatted(&W))
}

func (net RBFApproximator) Predict(value []float64) float64 {
	var v []float64
	for _, f := range net.HiddenLayer {
		v = append(v, f.Calculate(value))
	}
	/*vV := mat64.NewVector(len(net.HiddenLayer), v)
	wV := mat64.NewVector(len(net.Weights), net.Weights)
	result := mat64.Dot(vV, wV)
	return result*/
	return math.Abs(mat64.Dot(mat64.NewVector(len(net.HiddenLayer), v), mat64.NewVector(len(net.Weights), net.Weights)))
	//return mat64.Dot(mat64.NewVector(len(net.HiddenLayer), v), mat64.NewVector(len(net.Weights), net.Weights))
}

func (net RBFApproximator) PrintFunction() {
	fmt.Println("")
	for i, n := range net.HiddenLayer {
		fmt.Printf("%f*sqrt(%f*((x-%f)^2+(y-%f)^2))+", net.Weights[i], n.width, n.center[0], n.center[1])
		//fmt.Printf("%f*((e^(-((x-%f)^2+(y-%f)^2)))/(2*%f^2))+", net.Weights[i], n.center[0], n.center[1], n.width)
	}
}
func (net RBFApproximator) GetFunction() string {
	functionString := ""
	for i, n := range net.HiddenLayer {
		//functionString += fmt.Sprintf("%f*sqrt(%f*((x-%f)^2+(y-%f)^2))+", net.Weights[i], n.width, n.center[0], n.center[1])
		functionString += fmt.Sprintf("%f*(1-e^(-((x-%f)^2+(y-%f)^2)/(2*%f^2)))+", net.Weights[i], n.center[0], n.center[1], n.width)
		//functionString += fmt.Printf("%f*((e^(-((x-%f)^2+(y-%f)^2)))/(2*%f^2))+", net.Weights[i], n.center[0], n.center[1], n.width)
	}
	return functionString
}
