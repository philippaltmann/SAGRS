package Benchmark

import (
	"os"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
)

func FakeApproximator() {
	options := o.Options{
		Cycles:              20000,
		EvaluationRate:      2,
		SuggestToEvaluation: 8,
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          32,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4,
		WriteProgress:       true,
		Verbose:             true}

	//TODO track convergence if best fitness not changing
	//TODO adapt theshold to Suggested size
	fitnessFunction := Bohachevsky.EvaluateFitness
	Approximator := Approximation.GenerateEvaluation(fitnessFunction)

	_ = os.Mkdir("FakeApproximator", 0777)
	options.ProgressFileName = "FakeApproximator/Test"
	//initialPopulation := *new(Population.Population)
	a.Optimize(options, fitnessFunction, Approximator /*initialPopulation*/)
}
