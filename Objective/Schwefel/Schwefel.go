package Schwefel

import "math"

// Schwefel Empty Struct for Unversal Objective Application
type Schwefel struct{}

//Min returns the Objectives Lower Bound
func (s Schwefel) Min() float64 { return -100.0 }

//Max return the Objectives Upper Bound
func (s Schwefel) Max() float64 { return 100.0 }

//EvaluateFitness calculates the Objective fitness using the passed in value
func (s Schwefel) EvaluateFitness(values []float64) (fitness float64) {
	fitness = 418.9828872724339 * float64(len(values))
	for i := 0; i < len(values); i++ {
		xi := values[i]
		fitness -= xi * math.Sin(math.Sqrt(math.Abs(xi)))
	}
	return
}
