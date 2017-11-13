package H1

import "testing"

func TestEvaluation(t *testing.T) {
	testValue := []float64{8.6998, 6.7665}
	testResult := EvaluateFitness(testValue)

	if testResult < -0.0001 || testResult > 0.0001 {
		t.Errorf("Wrong Max at %f", testResult)
		t.Fail()
	}
}
