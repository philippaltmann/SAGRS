package Schaffer

import (
	"math"
)

const UpperBound float64 = 100
const LowerBound float64 = -100

func EvaluateFitness(values []float64) (fitness float64) {
	fitness = 0
	for i := 0; i < len(values)-1; i++ {
		xi := values[i]
		xj := values[i+1]
		fitness += math.Pow(
			(math.Pow(xi, 2)+math.Pow(xj, 2)), 0.25) *
			(math.Pow(math.Sin(
				(50*math.Pow(
					(math.Pow(xi, 2)+math.Pow(xj, 2)), 0.1))), 2) + 1)
	}
	return
}
