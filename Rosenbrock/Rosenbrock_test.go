package Rosenbrock

import (
	"math/rand"
	"testing"
)

func TestEvaluateFitness(t *testing.T) {
	//Testing if minimum is calculated as defined
	testValues := make([]float64, 32)
	for i := 0; i < len(testValues); i++ {
		testValues[i] = 1
	}
	testResult := EvaluateFitness(testValues)
	if testResult != 0 {
		t.Error("Minimum not 0\n")
		t.Error(testResult)
		t.Fail()
	}
}

func benchmarkRosenbrock(b *testing.B, dimensions int) {
	var vals []float64
	for v := 0; v < dimensions; v++ {
		vals = append(vals, rand.Float64()*200-100)
	}

	for n := 0; n < b.N; n++ {
		EvaluateFitness(vals)
	}
}
func BenchmarkRosenbrock4(b *testing.B)   { benchmarkRosenbrock(b, 4) }
func BenchmarkRosenbrock8(b *testing.B)   { benchmarkRosenbrock(b, 8) }
func BenchmarkRosenbrock16(b *testing.B)  { benchmarkRosenbrock(b, 16) }
func BenchmarkRosenbrock32(b *testing.B)  { benchmarkRosenbrock(b, 32) }
func BenchmarkRosenbrock64(b *testing.B)  { benchmarkRosenbrock(b, 64) }
func BenchmarkRosenbrock128(b *testing.B) { benchmarkRosenbrock(b, 128) }

/*
BenchmarkRosenbrock4-8     	 5000000	       278 ns/op
BenchmarkRosenbrock8-8     	 2000000	       624 ns/op
BenchmarkRosenbrock16-8    	 1000000	      1488 ns/op
BenchmarkRosenbrock32-8    	  500000	      3110 ns/op
BenchmarkRosenbrock64-8    	  200000	      6652 ns/op
BenchmarkRosenbrock128-8   	  100000	     13913 ns/op
*/
