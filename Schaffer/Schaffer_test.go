package Schaffer

import "testing"

func TestEvaluateFitness(t *testing.T) {
	//testValues := []float64{420.96874636, 420.96874636, 420.96874636, 420.96874636}
	testValues := []float64{420.96874636, 420.96874636}
	testResult := EvaluateFitness(testValues)
	//t.Error("Result: %f", testResult)

}
