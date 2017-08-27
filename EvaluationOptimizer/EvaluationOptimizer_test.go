package EvaluationOptimizer

import (
	"testing"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
)

func TestOptimizer(t *testing.T) {
	options := o.Options{
		Cycles:              10,
		EvaluationRate:      2,
		EvaluationPoolSize:  5,
		PopulationSize:      10,
		Dimensions:          2,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4}
	evaluatedProgress := Optimize(options, Bohachevsky.EvaluateFitness)
	if len(evaluatedProgress) != 10 {
		t.Error("Wrong Progress A\n")
		t.Error(evaluatedProgress)
		t.Fail()
	}
}
