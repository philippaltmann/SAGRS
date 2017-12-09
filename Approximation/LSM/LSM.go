package LSM

import (
	"github.com/gonum/matrix/mat64"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

//LSM Approximator for universal use with approximation interface
type LSM struct {
	ApproximationMatrix mat64.Dense
}

func Create() *LSM {
	return &LSM{}
}

//Update trains the Approximator to the given input and output
func (a *LSM) Update(population Population.Population) {

	y := generateFitnessVector(population)
	X := generateValueMatrix(population)

	var first mat64.Dense // construct a new zero-sized matrix
	first.Mul(X.T(), X)

	var firstInverse mat64.Dense
	firstInverse.Inverse(&first)

	var second mat64.Dense
	second.Mul(X.T(), y)

	var theta mat64.Dense
	theta.Mul(&firstInverse, &second)

	a.ApproximationMatrix = theta
}

//Predict applies the given value to the approximator
func (a *LSM) Predict(value []float64) float64 {
	valueVector := mat64.NewDense(1, len(value)+1, append([]float64{1.0}, value...))
	/*valueVector.Apply(func(_, _ int, v float64) float64 {
		return math.Pow(v, 2)
	}, valueVector)*/
	var resultMatrix mat64.Dense
	resultMatrix.Mul(valueVector, &a.ApproximationMatrix)
	result := resultMatrix.At(0, 0)
	//result := math.Abs(resultMatrix.At(0, 0))

	return result
}

func generateFitnessVector(population Population.Population) *mat64.Dense {
	n := len(population)
	var fitnessSlice []float64
	for i := 0; i < n; i++ {
		fitnessSlice = append(fitnessSlice, population[i].Fitness)
	}
	fitnessVector := mat64.NewDense(n, 1, fitnessSlice)
	return fitnessVector
}

func generateValueMatrix(population Population.Population) *mat64.Dense {
	n := len(population)          //Values
	k := len(population[0].Value) //Dimensions
	valueMatrix := mat64.NewDense(n, k+1, make([]float64, n*(k+1)))
	for i := 0; i < n; i++ {
		valueMatrix.SetRow(i, append([]float64{1.0}, population[i].Value...))
	}
	return valueMatrix
}
