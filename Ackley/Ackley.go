package Ackley

import "math"

const UpperBound float64 = 30.0
const LowerBound float64 = -15.0

func EvaluateFitness(values []float64) float64 {

	//20 - 20 exp (-0.2 sqrt(1/N sum(x_i^2))) +e - exp(1/N sum(cos(2pi x_i)))

	factor := 1.0 / float64(len(values))

	e1 := 20 * math.Exp(-0.2*math.Sqrt(factor*sum(values, func(v float64) float64 {
		return math.Pow(v, 2)
	})))

	e2 := math.Exp(factor * sum(values, func(v float64) float64 {
		return math.Cos(2 * math.Pi * v)
	}))

	return 20 - e1 + math.E - e2
}
func sum(values []float64, apply func(float64) float64) (sum float64) {
	for _, v := range values {
		sum += apply(v)
	}
	return
}
