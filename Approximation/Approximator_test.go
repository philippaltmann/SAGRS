package Approximation

import (
	"math"
	"testing"

	"github.com/philipp-altmann/SAGRS/Objective"
	"github.com/philipp-altmann/SAGRS/Population"
)

func TestApproximator(t *testing.T) {
	testPopulation := Population.InitRandomPopulation(10, 1, -10.0, 10.0)
	testObjective := Objective.GetObjective("")
	testPopulation.Evaluate(testObjective.EvaluateFitness)
	testRBF := GetApproximator("RBF")
	testRBF.Update(testPopulation)
	testLSM := GetApproximator("LSM")
	testLSM.Update(testPopulation)
	testEvaluation := GetApproximator("")
	testEvaluation.Update(testPopulation)

	for _, i := range testPopulation {
		//Test RBF
		if round(testRBF.Predict(i.Value), 2) != round(i.Fitness, 2) {
			t.Error("Failed to use RBF")
			t.Fail()
		}

		//Test LSM
		if round(testLSM.Predict(i.Value), 2) != round(i.Fitness, 2) {
			t.Error("Failed to use LSM")
			t.Fail()
		}

		//Test Evaluation
		if testEvaluation.Predict(i.Value) != i.Fitness {
			t.Error("Failed to use Evaluation")
			t.Fail()
		}
	}
}

func round(val float64, places int) (newVal float64) {
	roundOn := .5
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
