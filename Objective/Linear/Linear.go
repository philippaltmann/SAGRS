package Linear

// Linear Empty Struct for Unversal Objective Application
type Linear struct{}

//Min returns the Objectives Lower Bound
func (a Linear) Min() float64 { return -100.0 }

//Max return the Objectives Upper Bound
func (a Linear) Max() float64 { return 100.0 }

//EvaluateFitness calculates the Objective fitness using the passed in value
func (a Linear) EvaluateFitness(values []float64) float64 {
	return /*math.Abs(*/ sum(values, func(v float64) float64 {
		return v / float64(len(values))
	}) //)
}

func sum(values []float64, apply func(float64) float64) (sum float64) {
	for _, v := range values {
		sum += apply(v)
	}
	return
}
