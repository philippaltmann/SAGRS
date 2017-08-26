package Schwefel

import (
	"math"
)

const UpperBound float64 = 500
const LowerBound float64 = -500

func EvaluateFitness(values []float64) (fitness float64) {
	fitness = 418.9828872724339 * float64(len(values))
	for i := 0; i < len(values); i++ {
		xi := values[i]
		fitness -= xi * math.Sin(math.Sqrt(math.Abs(xi)))
	}
	return
}
