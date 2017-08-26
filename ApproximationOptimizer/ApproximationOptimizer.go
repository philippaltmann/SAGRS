package ApproximationOptimizer

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Generation"
	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

//TODO externalize
const cycles = 5000
const evaluationRate = 100

var evaluationPoolSize = 10

const populationSize = 100
const dimensions = 2

const selectionFactor = 0.6
const mutationFactor = 0.3
const recombinationFactor = 0.4

func Optimize() (progressEvaluated []float64, progressApproximated []float64) {

	cycle := 0

	//Init Evaluation Pool random
	/*var evaluationPool g.Generation
	for g := 0; g < evaluationPoolSize; g++ {
		evaluationPool = append(evaluationPool, i.GenerateRandomIndiviudal(dimensions))
		evaluationPool[g].EvaluateFitness()
	}*/
	const rangeSize = 200
	const rangeCenter = 0

	variations := int(math.Pow(2, dimensions))
	evaluationPoolSize = 0

	fmt.Print(variations)
	//Init Evaluation Pool systematically
	var evaluationPool g.Generation
	//Outer boundaries
	for d := 0; d < variations; d++ {
		var position []float64
		idx := d
		for di := 0; di < dimensions; di++ {
			factor := idx % 2
			baseValue := rangeSize/2 + rangeCenter
			position = append(position, float64(baseValue-factor*rangeSize))
			idx /= 2
		}
		evaluationPool = append(evaluationPool, i.GenerateIndividual(position))
		evaluationPoolSize++
	}

	var center []float64
	for i := 0; i < dimensions; i++ {
		center = append(center, float64(rangeCenter))
	}
	evaluationPool = append(evaluationPool, i.GenerateIndividual(center))
	evaluationPoolSize++

	fmt.Print("Init done")

	for g := 0; g < evaluationPoolSize; g++ {
		evaluationPool = append(evaluationPool, i.GenerateRandomIndiviudal(dimensions))
		evaluationPool[g].EvaluateFitness()
	}
	evaluationPool.Sort()

	//Init & Evaluate population
	var generation g.Generation
	for g := 0; g < populationSize; g++ {
		generation = append(generation, i.GenerateRandomIndiviudal(dimensions))
		generation[g].ApproximateFitness(evaluationPool)
	}

	generation.Sort()

	//Print best
	//generation.PrintBest(cycle)

	for cycle < cycles {

		//Evaluate best in generation & append to EvaluationPool
		if cycle%evaluationRate == 0 {
			best := i.GenerateIndividual(generation[0].Value)
			best.EvaluateFitness()

			for i := 0; i < evaluationRate; i++ {
				progressEvaluated = append(progressEvaluated, best.Fitness)
			}

			//TODO plot average fitness in evaluationPool
			fmt.Printf("Best approximated at %f, evaluated at %f\n", generation[0].Fitness, best.Fitness)
			evaluationPool = append(evaluationPool, best)
			evaluationPool.Sort()
			//evaluationPoolSize++
			//Cut worst to keep Evaluation Pool Size Constant
			evaluationPool = evaluationPool[:len(evaluationPool)-1]
			worstF := evaluationPool[len(evaluationPool)-1].Fitness
			bestF := evaluationPool[0].Fitness
			if worstF-bestF < 10 {
				//Check diversity by distances
				distances := 0.0
				for i := 0; i < evaluationPoolSize; i++ {
					for j := i + 1; j < evaluationPoolSize; j++ {
						distances += evaluationPool[i].DistanceTo(evaluationPool[j])
					}
				}
				//Calc mean distance devide by additions & dimensions
				distances /= 45
				distances /= 32
				fmt.Printf("Distance in Pool at %f\n", distances)
			}
			//Plotter.Plot2D(evaluationPool, int(cycle/evaluationRate))

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
			newIndiviudal.ApproximateFitness(evaluationPool)
			generation = append(generation, newIndiviudal)
		}

		//Evaluate
		for j := 0; j < populationSize; j++ {
			generation[j].ApproximateFitness(evaluationPool)
		}

		//Sort
		generation.Sort()

		//Append best to Hisory
		progressApproximated = append(progressApproximated, generation[0].Fitness)

		//Print best
		//generation.PrintBest(cycle)

		//End if good enough
		if generation[0].Fitness < 0.01 {
			//cycle = cycles
		}

		cycle++
	}

	return
}
