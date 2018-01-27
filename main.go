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
				rates := []int{0, 1, 2, 3, 4, 5, 6, 8, 16, 32, 64}
				approximator := args[2]
				objective := args[3]
				Benchmark.TestEvaluationRates(10, rates, approximator, objective)
			}
		}
	}
}
