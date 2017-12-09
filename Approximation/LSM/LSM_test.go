package LSM

import (
	"math"
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

var input [][]float64
var output []float64

var testSize = 10
var testDimensions = 1
var testPopulation = Population.InitRandomPopulation(testSize, testDimensions, -100.0, 100.0)

var testApproximator = LSM{}

func TestUpdate(t *testing.T) {
	testPopulation.Evaluate(func(v []float64) float64 {
		return v[0] + 1
	})
	testApproximator.Update(testPopulation)

	//Copy
	var factorCopy []float64
	factorCopy = mat64.Col(factorCopy, 0, &testApproximator.ApproximationMatrix)

	testPopulation.Evaluate(func(v []float64) float64 {
		return 2*v[0] - 1
	})
	testApproximator.Update(testPopulation)
	var factorUpdate []float64
	factorUpdate = mat64.Col(factorUpdate, 0, &testApproximator.ApproximationMatrix)
	//Compare
	for i, v := range factorCopy {
		if factorUpdate[i] == v {
			t.Error("Failed to Upadate Matrix")
			t.Fail()
		}
	}

	testPopulation.Evaluate(func(v []float64) float64 {
		return v[0] + 1
	})
	testApproximator.Update(testPopulation)
	factorUpdate = mat64.Col(factorUpdate, 0, &testApproximator.ApproximationMatrix)
	//Compare
	for i, v := range factorCopy {
		if round(factorUpdate[i], 5) != round(v, 5) {
			t.Error("Failed to Upadate Matrix")
			t.Fail()
		}
	}
}

func TestPredict(t *testing.T) {
	testPopulation.Evaluate(linearEvaluation)
	testApproximator.Update(testPopulation)

	for _, i := range testPopulation {
		if round(testApproximator.Predict(i.Value), 1) != round(i.Fitness, 1) {
			t.Error("Prediction Error")
			t.Fail()
		}
	}
}

func linearEvaluation(value []float64) float64 {
	return value[0] + 2
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
