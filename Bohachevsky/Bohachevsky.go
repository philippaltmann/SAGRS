package Bohachevsky

import "math"

const UpperBound float64 = 100
const LowerBound float64 = -100

func EvaluateFitness(values []float64) (fitness float64) {
	fitness = 0
	for i := 0; i < len(values)-1; i++ {
		xi := values[i]
		xj := values[i+1]
		fitness += (math.Pow(xi, 2) +
			2*math.Pow(xj, 2) -
			0.3*math.Cos(3*math.Pi*xi) -
			0.4*math.Cos(4*math.Pi*xj) +
			0.7)
	}
	return
}
