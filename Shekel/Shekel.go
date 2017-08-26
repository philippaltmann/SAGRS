package Shekel

import "math"

const UpperBound float64 = 100
const LowerBound float64 = -100

func EvaluateFitness(values []float64) (fitness float64) {
	const numberOfMaxima = 5

	A := [][]float64{{0.5, 0.25, 0.25, 0.75, 0.75}, {0.5, 0.25, 0.75, 0.25, 0.75}}
	c := []float64{0.002, 0.005, 0.005, 0.005, 0.005}

	//const numberOfMaxima = 1

	//A := [][]float64{{0.5}, {0.5}}
	//c := []float64{0.002}

	fitness = 0
	for i := 0; i < numberOfMaxima; i++ {
		fitness += c[i]
		var subsum float64 = 0
		for j := 0; j < len(values); j++ {
			subsum += math.Pow((values[j] - A[j][i]), 2)
		}
		fitness += subsum

	}
	fitness = 1 / fitness
	return
}
