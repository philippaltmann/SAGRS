package Ackley

import "math"

// Ackley Empty Struct for Unversal Objective Application
type Ackley struct{}

//Min returns the Objectives Lower Bound
func (a Ackley) Min() float64 { return -15.0 }

//Max return the Objectives Upper Bound
func (a Ackley) Max() float64 { return 30.0 }

//EvaluateFitness calculates the Objective fitness using the passed in value
func (a Ackley) EvaluateFitness(values []float64) float64 {
	factor := 1.0 / float64(len(values))
	e1 := 20 * math.Exp(
		-0.2*math.Sqrt(
			factor*sum(values, func(v float64) float64 {
				return math.Pow(v, 2)
			})))

	e2 := math.Exp(
		factor * sum(
			values, func(v float64) float64 {
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
