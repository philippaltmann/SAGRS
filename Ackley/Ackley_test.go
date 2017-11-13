package Ackley

import "testing"

func TestAckley(t *testing.T) {
	testValues := []float64{0.0, 0.0}
	testResult := EvaluateFitness(testValues)

	if testResult != 0.0 {
		t.Logf("False Minimum at %f", testResult)
		t.Fail()
	}
}
