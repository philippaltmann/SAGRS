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
			rates := []int{0, 1, 2, 3, 4, 5, 6, 8, 16, 32, 64}
			Benchmark.TestEvaluationRates(10, rates)
		}
	}
}
