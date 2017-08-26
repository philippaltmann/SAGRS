package ApproximationOptimizer

import (
	"fmt"
	"math/rand"
	"time"

	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

//TODO pass in fitness function
func Optimize(o o.Options) (progressEvaluated []float64, progressApproximated []float64) {
	cycle := 0

	//Init & Sort Evaluation Pool random
	var evaluationPool g.Population
	evaluationPool = g.InitRandomPopulation(o.EvaluationPoolSize, o.Dimensions)
	evaluationPool.Sort()

	ApproximationMatrix := a.GetLSMApproximator(evaluationPool)

	//Init & Sort population
	//TODO pass in approximation
	var population g.Population
	population = g.InitRandomPopulation(o.PopulationSize, o.Dimensions)
	population.Sort()

	//Optimize
	for cycle < o.Cycles {

		//Evaluate best in population & append to EvaluationPool
		if cycle%o.EvaluationRate == 0 {
			best := i.GenerateIndividual(population[0].Value)
			best.EvaluateFitness()

			for i := 0; i < o.EvaluationRate; i++ {
				progressEvaluated = append(progressEvaluated, best.Fitness)
			}

			//TODO plot average fitness in evaluationPool
			fmt.Printf("Best approximated at %f, evaluated at %f\n", population[0].Fitness, best.Fitness)
			evaluationPool = append(evaluationPool, best)
			evaluationPool.Sort()
			//o.EvaluationPoolSize++
			//Cut worst to keep Evaluation Pool Size Constant
			evaluationPool = evaluationPool[:len(evaluationPool)-1]
			worstF := evaluationPool[len(evaluationPool)-1].Fitness
			bestF := evaluationPool[0].Fitness
			if worstF-bestF < 10 {
				//Check diversity by distances
				distances := 0.0
				for i := 0; i < o.EvaluationPoolSize; i++ {
					for j := i + 1; j < o.EvaluationPoolSize; j++ {
						distances += evaluationPool[i].DistanceTo(evaluationPool[j])
					}
				}
				//Calc mean distance divide by additions & o.Dimensions
				distances /= 45
				distances /= 32
				fmt.Printf("Distance in Pool at %f\n", distances)
			}
			//Plotter.Plot2D(evaluationPool, int(cycle/o.EvaluationRate))

			//Update Approximator
			ApproximationMatrix = a.GetLSMApproximator(evaluationPool)

		}

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
			//newIndiviudal.ApproximateFitness(evaluationPool)
			population = append(population, newIndiviudal)
		}

		//Evaluate
		for j := 0; j < o.PopulationSize; j++ {
			population[j].Fitness = a.ApproximateFitness(population[j].Value, ApproximationMatrix)
			population[j].ApproximateFitness(evaluationPool)
		}

		//Sort
		population.Sort()

		//Append best to Hisory
		progressApproximated = append(progressApproximated, population[0].Fitness)

		//Print best
		//population.PrintBest(cycle)

		cycle++
	}

	return
}
