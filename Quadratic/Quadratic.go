package Quadratic

import "math"

const UpperBound float64 = 20.0
const LowerBound float64 = -20.0

func EvaluateFitness(values []float64) (fitness float64) {
	for _, v := range values {
		fitness += math.Pow(v, 2)
	}
	return
}
