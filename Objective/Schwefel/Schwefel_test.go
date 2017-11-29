package Schwefel

import (
	"math/rand"
	"testing"
)

func TestEvaluateFitness(t *testing.T) {
	//Testing if minimum is calculated as defined
	testValues := make([]float64, 32)
	for i := 0; i < len(testValues); i++ {
		testValues[i] = 420.96874636
	}
	testResult := EvaluateFitness(testValues)
	if testResult > 1e-10 || testResult < -1e-10 {
		t.Error("Minimum not 0\n")
		t.Error(testResult)
		t.Fail()
	}
}

func benchmarkSchwefel(b *testing.B, dimensions int) {
	var vals []float64
	for v := 0; v < dimensions; v++ {
		vals = append(vals, rand.Float64()*200-100)
	}

	for n := 0; n < b.N; n++ {
		EvaluateFitness(vals)
	}
}
func BenchmarkSchwefel4(b *testing.B)   { benchmarkSchwefel(b, 4) }
func BenchmarkSchwefel8(b *testing.B)   { benchmarkSchwefel(b, 8) }
func BenchmarkSchwefel16(b *testing.B)  { benchmarkSchwefel(b, 16) }
func BenchmarkSchwefel32(b *testing.B)  { benchmarkSchwefel(b, 32) }
func BenchmarkSchwefel64(b *testing.B)  { benchmarkSchwefel(b, 64) }
func BenchmarkSchwefel128(b *testing.B) { benchmarkSchwefel(b, 128) }

/*
BenchmarkSchwefel4-8     	20000000	        72.9 ns/op
BenchmarkSchwefel8-8     	10000000	       148 ns/op
BenchmarkSchwefel16-8    	 5000000	       290 ns/op
BenchmarkSchwefel32-8    	 3000000	       570 ns/op
BenchmarkSchwefel64-8    	 1000000	      1152 ns/op
BenchmarkSchwefel128-8   	 1000000	      2254 ns/op
*/
