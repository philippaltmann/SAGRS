package Approximation

import (
	"fmt"
	"math"

	"github.com/gonum/matrix/mat64"
	p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

type BiasFunction struct {
	center []float64
	width  float64
}

type RBFNet struct {
	HiddenLayer []BiasFunction
	Weights     []float64
}

func (f BiasFunction) Calculate(value []float64) float64 {
	return math.Exp(-(EuclideanDistance(value, f.center)) / (2 * math.Pow(f.width, 2)))
}

func EuclideanDistance(from []float64, to []float64) (distance float64) {
	for i, v := range from {
		distance += math.Pow((v - to[i]), 2)
	}
	return
}
func (net RBFNet) GetMatrix(input [][]float64) *mat64.Dense {
	var slice []float64
	for _, layer := range net.HiddenLayer {
		for _, value := range input {
			slice = append(slice, layer.Calculate(value))
		}
	}
	return mat64.NewDense(len(net.HiddenLayer), len(input), slice)
}

func GetRBFNet(input [][]float64, output []float64) RBFNet {
	net := new(RBFNet)

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
	functionMatrix := net.GetMatrix(input)
	functionMatrix.Inverse(functionMatrix)

	var W mat64.Dense
	W.Mul(functionMatrix, mat64.NewDense(len(output), 1, output))

	net.Weights = mat64.Col(net.Weights, 0, &W)
	fmt.Println(mat64.Formatted(&W))

	return *net
}

func (net RBFNet) Predict(value []float64) float64 {
	var v []float64
	for _, f := range net.HiddenLayer {
		v = append(v, f.Calculate(value))
	}
	return mat64.Dot(mat64.NewVector(len(net.HiddenLayer), v), mat64.NewVector(len(net.Weights), net.Weights))
}

var gamma = 10.0 // todo average distance * 2
var ppltn p.Population

func GetRBFApproximator(population p.Population) RBFNet {
	var input [][]float64
	var output []float64
	for _, i := range population {
		input = append(input, i.Value)
		output = append(output, i.Fitness)
	}
	return GetRBFNet(input, output)
}

func ApproximateRBF(approximationMatrix mat64.Dense, value []float64) (fitness float64) {
	size := len(ppltn)
	var p []float64
	for i := 0; i < size; i++ {
		p = append(p, math.Exp(-gamma*EuclideanDistance(value, ppltn[i].Value)))
	}
	P := mat64.NewVector(size, p)
	fitness = mat64.Dot(approximationMatrix.ColView(0), P)

	return
}
