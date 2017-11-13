package Quadratic

import "testing"

func TestEvaluate(t *testing.T) {
	best := 6.125
	maxValue := []float64{best}
	t.Log("Testing ")
	for d := 1; d < 32; d++ {
		if EvaluateFitness(maxValue) != 100 {
			t.Log(d)
			t.Error("Non matching max at " + string(d) + " dimensions\nFitness at ")
			t.Log(EvaluateFitness(maxValue))
			t.Fail()
		}
		maxValue = append(maxValue, best)
	}
	t.Log(" dimensions")
}
