package main

import (
	"fmt"
	"os"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Benchmark"
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
			if len(args) == 4 {
				//rates := []int{0, 1, 2, 3, 4, 5, 6, 8, 16, 32, 64}
				rates := []int{8, 32, 64}
				approximator := args[2]
				objective := args[3]
				Benchmark.TestEvaluationRates(10, rates, approximator, objective, 1000)
			} else {
				evaluationRateTestSuite()
			}
		}
	}
}

func evaluationRateTestSuite() {
	//reset = ["Reset", "NoReset"]
	approximators := []string{"LSM", "RBF"}
	objectives := []string{"Bohachevsky", "Ackley", "Schwefel"}
	rates := []int{0, 1, 2, 3, 4, 5, 6, 8, 16, 32, 64}
	cycles := 400
	for _, approximator := range approximators {
		for _, objective := range objectives {
			Benchmark.TestEvaluationRates(10, rates, approximator, objective, cycles)
		}
	}

}
