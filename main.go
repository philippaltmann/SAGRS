package main

import (
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
	e "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/EvaluationOptimizer"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Plotter"
)

func main() {
	options := o.Options{
		Cycles:              100000,
		EvaluationRate:      100,
		EvaluationPoolSize:  1000,
		PopulationSize:      100,
		Dimensions:          32,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4}
	approximatedProgressE, approximatedProgressA := a.Optimize(options, Bohachevsky.EvaluateFitness)
	evaluatedProgress := e.Optimize(options, Bohachevsky.EvaluateFitness)

	evaluatedPlottable := p.ConvertProgress(evaluatedProgress)
	approximatedPlottableE := p.ConvertProgress(approximatedProgressE)
	approximatedPlottableA := p.ConvertProgress(approximatedProgressA)

	p.Plot(evaluatedPlottable, approximatedPlottableE, approximatedPlottableA)

}
