package Approximation

import (
	"github.com/philipp-altmann/SAGRS/Approximation/Evaluation"
	"github.com/philipp-altmann/SAGRS/Approximation/LSM"
	"github.com/philipp-altmann/SAGRS/Approximation/RBF"
	"github.com/philipp-altmann/SAGRS/Population"
)

//Approximator interface for universal apporixmator usage
type Approximator interface {
	Update(Population.Population)
	Predict([]float64) float64
	Formatted() string
}

// GetApproximator generates a new Approximator
// creates Evaluating Approximator
// if no matching Approximator Model is specified
func GetApproximator(name string) Approximator {
	switch name {
	case "LSM":
		return LSM.Create()
	case "RBF":
		return RBF.Create()
	default:
		return Evaluation.Create(name)
	}
}
