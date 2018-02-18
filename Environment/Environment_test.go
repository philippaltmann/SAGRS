package Environment

import "testing"

func TestEnvironment(t *testing.T) {
	var testEnv = Environment{
		SuggestToEvaluation: 1,
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          2,
		SelectionFactor:     0.9,
		MutationFactor:      0.1,
		RecombinationFactor: 0.05,
		WriteProgress:       true,
		Verbose:             true,
		ResetPool:           true,
		Cycles:              1000,
		Approximator:        "RBF",
		Objective:           "Schwefel"}

	testEnv.Dump("Environment/test")
}
