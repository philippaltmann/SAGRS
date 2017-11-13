package Quadratic

import "math"

const UpperBound float64 = 20.0
const LowerBound float64 = -20.0

func EvaluateFitness(values []float64) (fitness float64) {
	center := -6.125
	fitness = 100
	for _, v := range values {
		fitness -= math.Pow(v+center, 2)
	}
	return
}
