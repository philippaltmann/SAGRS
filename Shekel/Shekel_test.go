package Shekel

import (
	"math/rand"
	"testing"
)

func TestEvaluateFitness(t *testing.T) {
	//Testing if minimum is calculated as defined
	testValues := make([]float64, 2)
	for i := 0; i < len(testValues); i++ {
		testValues[i] = 0.5
	}
	testResult := EvaluateFitness(testValues)
	if testResult < 1.5 {
		t.Error("Maximum not Met\n")
		t.Error(testResult)
		t.Fail()
	}
}

func BenchmarkShekel(b *testing.B) {
	var vals []float64
	dimensions := 2
	for v := 0; v < dimensions; v++ {
		vals = append(vals, rand.Float64()*200-100)
	}

	for n := 0; n < b.N; n++ {
		EvaluateFitness(vals)
	}
}

/*
BenchmarkShekel-8   	 5000000	       315 ns/op
*/
