package Evaluation

import (
	"github.com/philipp-altmann/SAGRS/Objective"
	"github.com/philipp-altmann/SAGRS/Population"
)

//Evaluation fake approximates, evaluating with a fitness function
type Evaluation struct {
	FitnessFunction func([]float64) float64
}

//Create new Evaluation instance
func Create(objective string) *Evaluation {
	return &Evaluation{Objective.GetObjective(objective).EvaluateFitness}
}

//Update trains the Approximator to the given input and output
func (a *Evaluation) Update(population Population.Population) {}

//Predict applies the given value to the approximator
func (a *Evaluation) Predict(value []float64) float64 {
	return a.FitnessFunction(value)
}

//Formatted converts current Meta Model to string
func (a *Evaluation) Formatted() string {
	return ""
}
