package Evaluation

import "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"

//Evaluation fake approximates, evaluating with a fitness function
type Evaluation struct {
	FitnessFunction func([]float64) float64
}

//Create new Evaluation instance
func Create(objective func([]float64) float64) *Evaluation {
	return &Evaluation{objective}
}

//Update trains the Approximator to the given input and output
func (a *Evaluation) Update(population Population.Population) {}

//Predict applies the given value to the approximator
func (a *Evaluation) Predict(value []float64) float64 {
	return a.FitnessFunction(value)
}
