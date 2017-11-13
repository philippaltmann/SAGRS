package main

import "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Benchmark"

func main() {

	//Benchmark.SimpleTest() //100, 7, 1
	Benchmark.TestEvaluationRates(100, 7, 0) //100, 7, 1
	//Benchmark.TestConvergence()

}
