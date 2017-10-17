package main

import (
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	e "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/EvaluationOptimizer"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Rosenbrock"
)

func FitnessTest(val []float64) float64 {
	fitness := 0.0
	for i := 0; i < len(val); i++ {
		fitness += float64(i+1) * val[i]
	}
	fitness += 5
	return fitness
}

func main() {
	options := o.Options{
		Cycles:              10000,
		EvaluationRate:      100,
		EvaluationPoolSize:  10,
		PopulationSize:      100,
		Dimensions:          20,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4}
	//approximatedProgressE, approximatedProgressA :=
	a.Optimize(options, Rosenbrock.EvaluateFitness)
	e.Optimize(options, Rosenbrock.EvaluateFitness)
	//p.WriteCSV(progressE, "evaluated")
	//p.WriteCSV(progressA, "approximated")

	//evaluatedProgress :=
	//e.Optimize(options, FitnessTest)

	/*evaluatedPlottable := p.ConvertProgress(evaluatedProgress)
	approximatedPlottableE := p.ConvertProgress(approximatedProgressE)
	approximatedPlottableA := p.ConvertProgress(approximatedProgressA)

	p.Plot(evaluatedPlottable, approximatedPlottableE, approximatedPlottableA)*/

}
