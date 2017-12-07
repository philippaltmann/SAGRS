package Bohachevsky

import "math"

// Bohachevsky Empty Struct for Unversal Objective Application
type Bohachevsky struct{}

//Min returns the Objectives Lower Bound
func (b Bohachevsky) Min() float64 { return -100.0 }

//Max return the Objectives Upper Bound
func (b Bohachevsky) Max() float64 { return 100.0 }

//EvaluateFitness calculates the Objective fitness using the passed in value
func (b Bohachevsky) EvaluateFitness(values []float64) (fitness float64) {
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
