package main

import (
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	e "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/EvaluationOptimizer"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Plotter"
)

func main() {
	options := o.Options{
		Cycles:              500,
		EvaluationRate:      100,
		EvaluationPoolSize:  10,
		PopulationSize:      100,
		Dimensions:          2,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4}
	approximatedProgressE, approximatedProgressA := a.Optimize(options)
	evaluatedProgress := e.Optimize(options)

	evaluatedPlottable := p.ConvertProgress(evaluatedProgress)
	approximatedPlottableE := p.ConvertProgress(approximatedProgressE)
	approximatedPlottableA := p.ConvertProgress(approximatedProgressA)

	p.Plot(evaluatedPlottable, approximatedPlottableE, approximatedPlottableA)

}
