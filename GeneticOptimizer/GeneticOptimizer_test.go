package GeneticOptimizer

import (
	"testing"

	"github.com/philipp-altmann/SAGRS/Environment"
)

func TestOptimizer(t *testing.T) {
	testEnvironment := Environment.Environment{
		PopulationSize:      10,
		Cycles:              2,
		Dimensions:          2,
		SelectionFactor:     0.9,
		MutationFactor:      0.1,
		RecombinationFactor: 0.05,
		WriteProgress:       false,
		Verbose:             false,
		Approximator:        "Linear",
		Objective:           "Linear"}
	Optimize(testEnvironment)
}
