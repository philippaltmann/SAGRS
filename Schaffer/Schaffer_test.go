package Schaffer

import (
	"math/rand"
	"testing"
)

func TestEvaluateFitness(t *testing.T) {
	//Testing if minimum is calculated as defined
	testValues := make([]float64, 32)
	testResult := EvaluateFitness(testValues)
	if testResult != 0 {
		t.Error("Minimum not 0\n")
		t.Fail()
	}
}

func benchmarkSchaffer(b *testing.B, dimensions int) {
	var vals []float64
	for v := 0; v < dimensions; v++ {
		vals = append(vals, rand.Float64()*200-100)
	}

	for n := 0; n < b.N; n++ {
		EvaluateFitness(vals)
	}
}
func BenchmarkSchaffer4(b *testing.B)   { benchmarkSchaffer(b, 4) }
func BenchmarkSchaffer8(b *testing.B)   { benchmarkSchaffer(b, 8) }
func BenchmarkSchaffer16(b *testing.B)  { benchmarkSchaffer(b, 16) }
func BenchmarkSchaffer32(b *testing.B)  { benchmarkSchaffer(b, 32) }
func BenchmarkSchaffer64(b *testing.B)  { benchmarkSchaffer(b, 64) }
func BenchmarkSchaffer128(b *testing.B) { benchmarkSchaffer(b, 128) }

/*
BenchmarkSchaffer4-8     	 1000000	      1079 ns/op
BenchmarkSchaffer8-8     	  500000	      2608 ns/op
BenchmarkSchaffer16-8    	  300000	      5815 ns/op
BenchmarkSchaffer32-8    	  100000	     12170 ns/op
BenchmarkSchaffer64-8    	   50000	     25785 ns/op
BenchmarkSchaffer128-8   	   30000	     52123 ns/op
*/
