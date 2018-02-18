package Linear

import (
	"math/rand"
	"testing"
)

func TestEvaluateFitness(t *testing.T) {
	//Testing if minimum is calculated as defined
	testLinear := Linear{}
	testLinear.Min()
	testLinear.Max()
	testValues := make([]float64, 32)
	testResult := testLinear.EvaluateFitness(testValues)
	if testResult != 0 {
		t.Error("Minimum not 0\n")
		t.Error(testResult)
		t.Fail()
	}
}

func benchmarkLinear(b *testing.B, dimensions int) {
	testLinear := Linear{}
	var vals []float64
	for v := 0; v < dimensions; v++ {
		vals = append(vals, rand.Float64()*200-100)
	}

	for n := 0; n < b.N; n++ {
		testLinear.EvaluateFitness(vals)
	}
}
func BenchmarkLinear4(b *testing.B)   { benchmarkLinear(b, 4) }
func BenchmarkLinear8(b *testing.B)   { benchmarkLinear(b, 8) }
func BenchmarkLinear16(b *testing.B)  { benchmarkLinear(b, 16) }
func BenchmarkLinear32(b *testing.B)  { benchmarkLinear(b, 32) }
func BenchmarkLinear64(b *testing.B)  { benchmarkLinear(b, 64) }
func BenchmarkLinear128(b *testing.B) { benchmarkLinear(b, 128) }

/*
BenchmarkLinear4-8     	100000000	        17.6 ns/op
BenchmarkLinear8-8     	 50000000	        28.4 ns/op
BenchmarkLinear16-8    	 30000000	        50.4 ns/op
BenchmarkLinear32-8    	 20000000	       103.0 ns/op
BenchmarkLinear64-8    	 10000000	       186.0 ns/op
BenchmarkLinear128-8   	  5000000	       377.0 ns/op
*/
