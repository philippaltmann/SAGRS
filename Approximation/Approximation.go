package Approximation

import (
	"github.com/gonum/matrix/mat64"
	p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func generateFitnessVector(population p.Population) *mat64.Dense {
	var fitnessSlice []float64
	for _, individual := range population {
		fitnessSlice = append(fitnessSlice, individual.Fitness)
	}
	fitnessVector := mat64.NewDense(len(population), 1, fitnessSlice)
	return fitnessVector
}

func generateValueMatrix(population p.Population) *mat64.Dense {
	var valueSlice []float64
	for _, individual := range population {
		for _, value := range individual.Value {
			valueSlice = append(valueSlice, value)
		}
	}
	valueMatrix := mat64.NewDense(len(population), len(valueSlice)/len(population), valueSlice)
	return valueMatrix
}

func GetLSMApproximator(population p.Population) (approxmationMatrix mat64.Dense) {

	y := generateFitnessVector(population)
	X := generateValueMatrix(population)
	XT := X.T()
	var first mat64.Dense // construct a new zero-sized matrix
	first.Mul(XT, X)

	//r, c := first.Dims()
	var firstInverse mat64.Dense
	firstInverse.Inverse(&first)

	var second mat64.Dense
	second.Mul(XT, y)

	var theta mat64.Dense
	theta.Mul(&firstInverse, &second)

	return theta
}

func ApproximateFitness(value []float64, ApproximationMatrix mat64.Dense) (fitness float64) {
	valueVector := mat64.NewDense(1, len(value), value)
	var result mat64.Dense
	result.Mul(valueVector, &ApproximationMatrix)
	return result.At(0, 0)
}
