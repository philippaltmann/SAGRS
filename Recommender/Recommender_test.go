package Recommender

import (
	"testing"

	e "github.com/philipp-altmann/SAGRS/Environment"
)

func TestOptimizer(t *testing.T) {
	testEnvironment := e.Environment{
		EvaluationPoolSize:  10,
		PopulationSize:      10,
		Dimensions:          2,
		SelectionFactor:     0.9,
		MutationFactor:      0.1,
		RecombinationFactor: 0.05,
		WriteProgress:       false,
		Verbose:             false,
		ResetPool:           true,
		EvaluationRate:      1,
		Cycles:              10,
		SuggestToEvaluation: 1,
		Approximator:        "LSM",
		Objective:           "Bohachevsky"}

	Run(testEnvironment)
}
