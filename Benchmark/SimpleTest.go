package Benchmark

import (
	"fmt"
	"os"

	"github.com/buger/goterm"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Ackley"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
)

func SimpleTest() {
	options := o.Options{
		Cycles:               1000,
		SuggestToEvaluation:  1,
		EvaluationRate:       1,
		EvaluationPoolSize:   2,
		PopulationSize:       100,
		Dimensions:           2,
		SelectionFactor:      0.6,
		MutationFactor:       0.3,
		RecombinationFactor:  0.4,
		WriteProgress:        false,
		Verbose:              true,
		ResetPool:            true,
		ConvergenceThreshold: 100}

	fitnessFunction := Ackley.EvaluateFitness
	Approximator := new(Approximation.RBFApproximator)

	var functionFile *os.File

	goterm.Clear()
	functionFile, _ = os.Create("Functions.txt")
	defer functionFile.Close()
	//options.ProgressFileName = "EvaluationRateTest" + aType + "/EvaluationRateTest" + strconv.Itoa(options.EvaluationRate) + "_" + strconv.Itoa(i)
	_, cycle := a.Optimize(options, fitnessFunction, Approximator)
	fmt.Printf("Converged in cycle %d", cycle)
	functionFile.WriteString(Approximator.GetFunction() + "\n\n")
}
