package main

import (
	"fmt"
	"os"

	"github.com/philipp-altmann/SAGRS/Benchmark"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Please Specify What you want to run (e.g. benchmark rates)")
		return
	}
	switch args[0] {
	case "benchmark":
		switch args[1] {
		case "rates":
			evaluationRateTestSuite()
			break
		case "suggestions":
			suggestionsTestSuite()
			break
		case "compare":
			comparisonTestSuite()
			break
		}
	}
}

func evaluationRateTestSuite() {
	//reset = ["Reset", "NoReset"]
	approximators := []string{"LSM", "RBF"}
	objectives := []string{"Bohachevsky", "Ackley", "Schwefel"}
	//rates := []int{0, 1, 2, 3, 4, 5, 6, 8, 16, 32, 64}
	rates := []int{1, 2, 4, 8, 16, 32, 64}
	cycles := 1000
	for _, approximator := range approximators {
		for _, objective := range objectives {
			Benchmark.TestEvaluationRates(10, rates, approximator, objective, cycles)
		}
	}
}

func suggestionsTestSuite() {
	//reset = ["Reset", "NoReset"]
	//Rates for    LSMB, LSMA, LSMS, RBFB, RBFA, RBFS
	rates := []int{64, 64, 1, 4, 8, 1} //[]int{64, 4, 8, 1, 4, 8}
	resets := []bool{false, true, true, false, false, true}
	approximators := []string{"LSM", "RBF"}
	objectives := []string{"Bohachevsky", "Ackley", "Schwefel"}
	suggestions := []int{1, 2, 3, 4, 5, 6, 7, 8}
	cycles := 1000
	for i, approximator := range approximators {
		for j, objective := range objectives {
			rate := rates[3*i+j]
			reset := resets[3*i+j]
			Benchmark.TestSuggestions(10, rate, reset, suggestions, approximator, objective, cycles)
		}
	}
}

func comparisonTestSuite() {
	//LSM Bohachevsky
	//Compare, Rate, Reset, Suggestion, Approximator, Objective, Cycles
	Benchmark.Compare(10, 64, false, 4, "LSM", "Bohachevsky", 100)

	//RBF Bohachevsky
	//Compare, Rate, Reset, Suggestion, Approximator, Objective, Cycles
	Benchmark.Compare(10, 4, false, 8, "RBF", "Bohachevsky", 100)

	//LSM Ackley
	//Compare, Rate, Reset, Suggestion, Approximator, Objective, Cycles
	Benchmark.Compare(10, 64, true, 7, "LSM", "Ackley", 100)

	//RBF Ackley
	//Compare, Rate, Reset, Suggestion, Approximator, Objective, Cycles
	Benchmark.Compare(10, 8, false, 7, "RBF", "Ackley", 120)

	//LSM Schwefel
	//Compare, Rate, Reset, Suggestion, Approximator, Objective, Cycles
	Benchmark.Compare(10, 1, true, 6, "LSM", "Schwefel", 100)

	//RBF Schwefel
	//Compare, Rate, Reset, Suggestion, Approximator, Objective, Cycles
	Benchmark.Compare(10, 1, true, 8, "RBF", "Schwefel", 100)
}
