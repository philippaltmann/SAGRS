package EvaluationOptimizer

import (
	"math"
	"math/rand"
	"time"

	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Generation"
	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

//TODO externalize
const cycles = 5000
const populationSize = 100
const dimensions = 32

const selectionFactor = 0.6
const mutationFactor = 0.3
const recombinationFactor = 0.4

func Optimize() (progress []float64) {

	cycle := 0
	bestFitness := math.MaxFloat64

	var generation g.Generation
	//Init & Evaluate population
	for g := 0; g < populationSize; g++ {
		generation = append(generation, i.GenerateRandomIndiviudal(dimensions))
		generation[g].EvaluateFitness()
	}

	generation.Sort()
	bestFitness = generation[0].Fitness
	progress = append(progress, bestFitness)

	//Print best
	generation.PrintBest(cycle)

	for cycle < cycles {

		if cycle%100 == 0 {
			//Plotter.Plot2D(generation, int(cycle/100))

		}

		//Select
		generation = generation[:populationSize-populationSize*selectionFactor]

		//Mutate
		/*Start incrementing from 1 to prevent mutation of best Individual*/
		for p := 1; p < len(generation); p++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mutate := r.Float32()

			if mutate > 1-mutationFactor {
				generation[p].Mutate()
			}

		}
		//Recombine
		for p := 0; p < len(generation); p++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			recombine := r.Float32()

			if recombine > 1-recombinationFactor {
				combineWith := r.Intn(len(generation))
				newIndividual := generation[p].Recombine(generation[combineWith])
				if len(generation) < populationSize {
					generation = append(generation, newIndividual)
				}
			}

		}

		//Fillup
		for len(generation) < populationSize {
			newIndiviudal := i.GenerateRandomIndiviudal(dimensions)
			newIndiviudal.EvaluateFitness()
			generation = append(generation, newIndiviudal)
		}

		//Evaluate
		for j := 0; j < populationSize; j++ {
			generation[j].EvaluateFitness()
		}

		//Sort
		generation.Sort()

		//Append best to Hisory
		progress = append(progress, generation[0].Fitness)

		//Print best
		generation.PrintBest(cycle)

		cycle++
	}

	return
}
