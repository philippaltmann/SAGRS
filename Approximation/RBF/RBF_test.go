package RBF

import (
	"math"
	"testing"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Objective"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

var input = [][]float64{{-2}, {-1}, {0}, {1}, {2}}
var output = []float64{4, 1, 0, 1, 4}

var testSize = 10
var testDimensions = 1
var testPopulation = Population.InitRandomPopulation(testSize, testDimensions, -100.0, 100.0)

var testApproximator = RBF{}
var testObjective = Objective.GetObjective("Ackley")

func TestUpdate(t *testing.T) {
	testPopulation.Evaluate(testObjective.EvaluateFitness)
	testApproximator.Update(testPopulation)

	//Copy
	input, output = getInputOutput(testPopulation)
	weights := make([]float64, len(testApproximator.Weights))
	copy(weights, testApproximator.Weights)
	width := make([]float64, len(testApproximator.HiddenLayer))
	for i := range width {
		width[i] = testApproximator.HiddenLayer[i].width
	}
	centers := make([][]float64, len(testApproximator.HiddenLayer))
	for i := range centers {
		centers[i] = make([]float64, len(testApproximator.HiddenLayer[i].center))
		copy(centers[i], testApproximator.HiddenLayer[i].center)
	}

	testPopulation.Mutate(1)
	testPopulation.Evaluate(testObjective.EvaluateFitness)
	testApproximator.Update(testPopulation)

	//Compare
	input, output = getInputOutput(testPopulation)
	for i, w := range weights {
		if w == testApproximator.Weights[i] {
			t.Error("Failed to update weights")
			t.Fail()
		}
	}
	for i := range width {
		if width[i] == testApproximator.HiddenLayer[i].width {
			t.Error("Failed to update width")
			t.Fail()
		}
	}
	for i, c := range centers {
		for j := range c {
			if c[j] == testApproximator.HiddenLayer[i].center[j] {
				//First is never muated
				if i > 0 {
					t.Error("Failed to update centers")
					t.Fail()
				}
			}
		}
	}
}

func TestPredict(t *testing.T) {
	testApproximator.Update(testPopulation)
	input, output = getInputOutput(testPopulation)
	for i, v := range input {
		if round(testApproximator.Predict(v), 1) != round(output[i], 1) {
			t.Error("Prediction Error")
			t.Log(testApproximator.Predict(v))
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

func getInputOutput(population Population.Population) (input [][]float64, output []float64) {
	for _, i := range population {
		input = append(input, i.Value)
		output = append(output, i.Fitness)
	}
	return
}
