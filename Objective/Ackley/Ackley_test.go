package Ackley

import "testing"

func TestAckley(t *testing.T) {
	testAckley := Ackley{}
	testAckley.Min()
	testAckley.Max()
	testValues := []float64{0.0, 0.0}
	testResult := testAckley.EvaluateFitness(testValues)

	if testResult != 0.0 {
		t.Logf("False Minimum at %f", testResult)
		t.Fail()
	}
}
