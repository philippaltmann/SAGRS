package Approximation

import (
	"testing"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func TestRBFApproximator(t *testing.T) {
	size := 100
	dimensions := 2

	//Init Approximator
	//testApproximator := GenerateRBFApproximator(size, dimensions)
	testApproximator := new(RBFApproximator)

	/*wrongHiddenLayerSize := len(testApproximator.HiddenLayer) != size
	wrongWeightsSize := len(testApproximator.Weights) != size
	wrongDimensions := len(testApproximator.HiddenLayer[0].center) != dimensions
	t.Log("Testing Size & Dimension")
	if wrongHiddenLayerSize || wrongWeightsSize || wrongDimensions {
		t.Error("Wrong Size")
		t.Fail()
	}*/
	//var a Approximator
	testPopulation := Population.InitRandomPopulation(size, dimensions)
	testPopulation.Evaluate(Bohachevsky.EvaluateFitness)

	testApproximator.Update(testPopulation)
	testPrediction := testApproximator.Predict([]float64{0, 0})
	if testPrediction < 1.0 {
		t.Log("Success")
	}
}
func TestLSMApproximator(t *testing.T) {
	size := 100
	dimensions := 2

	//Init Approximator
	testApproximator := new(LSMApproximator)

	//var a Approximator
	testPopulation := Population.InitRandomPopulation(size, dimensions)
	testPopulation.Evaluate(Bohachevsky.EvaluateFitness)

	testApproximator.Update(testPopulation)

	t.Log(testApproximator)

	testPopulation2 := Population.InitRandomPopulation(size, dimensions)
	testPopulation2.Evaluate(Bohachevsky.EvaluateFitness)
	testApproximator.Update(testPopulation2)

	t.Log(testApproximator)
}
