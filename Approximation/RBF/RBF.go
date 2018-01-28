package RBF

import (
	"math"

	"github.com/gonum/matrix/mat64"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

//RBF Approximator for universal use with approximation interface
type RBF struct {
	HiddenLayer []BiasFunction
	Weights     []float64
}

//Create new instance of RBF
func Create() *RBF {
	return &RBF{}
}

//getFunctionMatrix applies the hidden layer bias functions to the input
func (a RBF) getFunctionMatrix(input [][]float64) *mat64.Dense {
	var functionSlice []float64
	for _, layer := range a.HiddenLayer {
		for _, value := range input {
			functionSlice = append(functionSlice, layer.Calculate(value))
		}
	}
	return mat64.NewDense(len(a.HiddenLayer), len(input), functionSlice)
}

//Update trains the Approximator to the given input and output
func (a *RBF) Update(population Population.Population) {
	var input [][]float64
	var output []float64
	for _, i := range population {
		input = append(input, i.Value)
		output = append(output, i.Fitness)
	}

	var averageDistance float64
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			averageDistance += EuclideanDistance(input[i], input[j])
		}
	}
	averageDistance = math.Sqrt(averageDistance)
	averageDistance /= float64(len(input))

	//"Train" Layers
	a.HiddenLayer = make([]BiasFunction, len(input))
	for i := 0; i < len(input); i++ {
		a.HiddenLayer[i].center = append([]float64(nil), input[i]...)
		a.HiddenLayer[i].width = averageDistance
	}

	//"Train" Weights
	functionMatrix := a.getFunctionMatrix(input)
	functionMatrix.Inverse(functionMatrix)

	var W mat64.Dense
	W.Mul(functionMatrix, mat64.NewDense(len(output), 1, output))

	a.Weights = mat64.Col(a.Weights, 0, &W)
}

//Predict applies the given value to the approximator
func (a RBF) Predict(value []float64) float64 {
	var v []float64
	for _, f := range a.HiddenLayer {
		v = append(v, f.Calculate(value))
	}
	valueVector := mat64.NewVector(len(a.HiddenLayer), v)
	weightVector := mat64.NewVector(len(a.Weights), a.Weights)
	result := mat64.Dot(valueVector, weightVector)
	return result
}

func (a *RBF) Formatted() string {
	return ""
}
