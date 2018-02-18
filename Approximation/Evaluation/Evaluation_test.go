package Evaluation

import (
	"testing"

	"github.com/philipp-altmann/SAGRS/Objective"
	"github.com/philipp-altmann/SAGRS/Population"
)

var testObjective = Objective.GetObjective("Ackley")
var testApproximator = Create("Ackley")
var testPopulation = Population.InitRandomPopulation(10, 2, -10.0, 10.0)

func TestEvaluation(t *testing.T) {
	testApproximator.Formatted()
	testPopulation.Evaluate(testApproximator.Predict)
	for _, i := range testPopulation {
		if i.Fitness != testObjective.EvaluateFitness(i.Value) {
			t.Error("Failed to Use Objective")
			t.Fail()
		}
	}
}
