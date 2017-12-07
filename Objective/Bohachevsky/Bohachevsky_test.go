package Bohachevsky

import (
	"math/rand"
	"testing"
)

func TestEvaluateFitness(t *testing.T) {
	//Testing if minimum is calculated as defined
	testBohachevsky := Bohachevsky{}
	testBohachevsky.Min()
	testBohachevsky.Max()
	testValues := make([]float64, 32)
	testResult := testBohachevsky.EvaluateFitness(testValues)
	if testResult != 0 {
		t.Error("Minimum not 0\n")
		t.Error(testResult)
		t.Fail()
	}
}

func benchmarkBohachevsky(b *testing.B, dimensions int) {
	testBohachevsky := Bohachevsky{}
	var vals []float64
	for v := 0; v < dimensions; v++ {
		vals = append(vals, rand.Float64()*200-100)
	}

	for n := 0; n < b.N; n++ {
		testBohachevsky.EvaluateFitness(vals)
	}
}
func BenchmarkBohachevsky4(b *testing.B)   { benchmarkBohachevsky(b, 4) }
func BenchmarkBohachevsky8(b *testing.B)   { benchmarkBohachevsky(b, 8) }
func BenchmarkBohachevsky16(b *testing.B)  { benchmarkBohachevsky(b, 16) }
func BenchmarkBohachevsky32(b *testing.B)  { benchmarkBohachevsky(b, 32) }
func BenchmarkBohachevsky64(b *testing.B)  { benchmarkBohachevsky(b, 64) }
func BenchmarkBohachevsky128(b *testing.B) { benchmarkBohachevsky(b, 128) }

/*
BenchmarkBohachevsky4-8     	 5000000	       264 ns/op
BenchmarkBohachevsky8-8     	 2000000	       609 ns/op
BenchmarkBohachevsky16-8    	 1000000	      1288 ns/op
BenchmarkBohachevsky32-8    	  500000	      2858 ns/op
BenchmarkBohachevsky64-8    	  300000	      5858 ns/op
BenchmarkBohachevsky128-8   	  100000	     12959 ns/op
*/
