package Rosenbrock

import (
	"math"
)

const UpperBound float64 = math.MaxFloat64
const LowerBound float64 = -math.MaxFloat64

func EvaluateFitness(values []float64) (fitness float64) {
	fitness = 0
	for i := 0; i < len(values)-1; i++ {
		xi := values[i]
		xj := values[i+1]
		fitness += math.Pow(
			1-xi, 2) + 100*math.Pow(
			(xj-math.Pow(xi, 2)), 2)
	}
	return
}
