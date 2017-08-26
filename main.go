package main

import e "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/EvaluationOptimizer"
import a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
import p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Plotter"

func main() {
	approximatedProgressE, approximatedProgressA := a.Optimize()
	evaluatedProgress := e.Optimize()

	evaluatedPlottable := p.ConvertProgress(evaluatedProgress)
	approximatedPlottableE := p.ConvertProgress(approximatedProgressE)
	approximatedPlottableA := p.ConvertProgress(approximatedProgressA)

	p.Plot(evaluatedPlottable, approximatedPlottableE, approximatedPlottableA)

}
