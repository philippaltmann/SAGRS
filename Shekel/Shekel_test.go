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
BenchmarkShekel4-8     	20000000	        72.9 ns/op
BenchmarkShekel8-8     	10000000	       148 ns/op
BenchmarkShekel16-8    	 5000000	       290 ns/op
BenchmarkShekel32-8    	 3000000	       570 ns/op
BenchmarkShekel64-8    	 1000000	      1152 ns/op
BenchmarkShekel128-8   	 1000000	      2254 ns/op
*/
