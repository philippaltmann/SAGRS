package EvaluationOptimizer

import (
	"math"
	"math/rand"
	"time"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func Optimize(o o.Options) (progress []float64) {

	cycle := 0
	bestFitness := math.MaxFloat64

	//Init & Sort population
	var population g.Population
	population = g.InitRandomPopulation(o.PopulationSize, o.Dimensions)
	population.Sort()

	bestFitness = population[0].Fitness
	progress = append(progress, bestFitness)

	//Print best
	population.PrintBest(cycle)

	for cycle < o.Cycles {

		//Select
		newSize := o.PopulationSize - int(float64(o.PopulationSize)*o.SelectionFactor)
		population = population[:newSize]

		//Mutate
		/*Start incrementing from 1 to prevent mutation of best Individual*/
		for p := 1; p < len(population); p++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mutate := r.Float64()

			if mutate > 1-o.MutationFactor {
				population[p].Mutate()
			}

		}
		//Recombine
		for p := 0; p < len(population); p++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			recombine := r.Float64()

			if recombine > 1-o.RecombinationFactor {
				combineWith := r.Intn(len(population))
				newIndividual := population[p].Recombine(population[combineWith])
				if len(population) < o.PopulationSize {
					population = append(population, newIndividual)
				}
			}

		}

		//Fillup
		for len(population) < o.PopulationSize {
			newIndiviudal := i.GenerateRandomIndiviudal(o.Dimensions)
			newIndiviudal.EvaluateFitness()
			population = append(population, newIndiviudal)
		}

		//Evaluate
		for j := 0; j < o.PopulationSize; j++ {
			population[j].EvaluateFitness()
		}

		//Sort
		population.Sort()

		//Append best to Hisory
		progress = append(progress, population[0].Fitness)

		//Print best
		population.PrintBest(cycle)

		cycle++
	}

	return
}
