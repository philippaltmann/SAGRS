package Evaluation

import (
	"testing"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Objective"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

var testObjective = Objective.GetObjective("Ackley")
var testApproximator = Create(testObjective.EvaluateFitness)
var testPopulation = Population.InitRandomPopulation(10, 2, -10.0, 10.0)

func TestEvaluation(t *testing.T) {
	testPopulation.Evaluate(testApproximator.Predict)
	for _, i := range testPopulation {
		if i.Fitness != testObjective.EvaluateFitness(i.Value) {
			t.Error("Failed to Use Objective")
			t.Fail()
		}
	}
}
