package EvaluationOptimizer

import (
	"fmt"

	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func Optimize(o o.Options, FitnessFunction func([]float64) float64) (progress []float64) {

	cycle := 0

	//Init, Approximate & Sort Population
	population := g.InitRandomPopulation(o.PopulationSize, o.Dimensions)
	population.Evaluate(FitnessFunction)
	population.Sort()

	//Print best
	population.PrintBest(cycle)

	for cycle < o.Cycles {

		//Select
		newSize := o.PopulationSize - int(float64(o.PopulationSize)*o.SelectionFactor)
		population = population[:newSize]

		//Mutate
		population.Mutate(o.MutationFactor)

		//Recombine
		population.Recombine(o.RecombinationFactor, o.PopulationSize)

		//Fillup
		population.Fillup(o.PopulationSize, o.Dimensions)

		//Evaluate using Fitness Function
		population.Evaluate(FitnessFunction)

		//Sort
		population.Sort()

		//Append best to Hisory
		progress = append(progress, population[0].Fitness)

		//Print best
		//population.PrintBest(cycle)
		fmt.Printf("Cycle %d: Best evaluated at %f\n", cycle, population[0].Fitness)

		cycle++
	}

	return
}
