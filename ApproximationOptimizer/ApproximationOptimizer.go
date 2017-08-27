package ApproximationOptimizer

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

var ApproximationMatrix mat64.Dense

func Approximator(value []float64) float64 {
	return a.ApproximateFitness(value, ApproximationMatrix)
}

func Optimize(o o.Options, FitnessFunction func([]float64) float64) (progressEvaluated []float64, progressApproximated []float64) {
	cycle := 0

	//Init, Evaluate & Sort Evaluation Pool random
	evaluationPool := g.InitRandomPopulation(o.EvaluationPoolSize, o.Dimensions)
	evaluationPool.Evaluate(FitnessFunction)
	evaluationPool.Sort()

	//Calculate Approximation Matrix for LSM Approximation
	ApproximationMatrix = a.GetLSMApproximator(evaluationPool)

	//Init, Approximate & Sort Population
	population := g.InitRandomPopulation(o.PopulationSize, o.Dimensions)
	population.Evaluate(Approximator)
	population.Sort()

	//Optimize
	for cycle < o.Cycles {

		//TODO plot average fitness in evaluationPool

		//Evaluate best in population & append to EvaluationPool
		if cycle%o.EvaluationRate == 0 {
			best := i.GenerateIndividual(population[0].Value)
			best.Fitness = FitnessFunction(best.Value)

			for i := 0; i < o.EvaluationRate; i++ {
				progressEvaluated = append(progressEvaluated, best.Fitness)
			}

			evaluationPool = append(evaluationPool, best)
			evaluationPool.Sort()

			//Cut worst to keep Evaluation Pool Size Constant
			evaluationPool = evaluationPool[:len(evaluationPool)-1]

			//Update Approximator
			ApproximationMatrix = a.GetLSMApproximator(evaluationPool)

			/*var distance float64
			for i := 0; i < len(evaluationPool); i++ {
				for j := i; j < len(evaluationPool); j++ {
					distance += a.DistanceTo(evaluationPool[i].Value, evaluationPool[j].Value)
				}
			}
			distance /= float64(len(evaluationPool))
			distance /= float64(len(evaluationPool))*/

			//Print Best
			fmt.Printf("Best approximated at %f, evaluated at %f\n\tBest in pool at %f\n", population[0].Fitness, best.Fitness, evaluationPool[0].Fitness)
		}

		//Select
		newSize := o.PopulationSize - int(float64(o.PopulationSize)*o.SelectionFactor)
		population = population[:newSize]

		//Mutate
		population.Mutate(o.MutationFactor)

		//Recombine
		population.Recombine(o.RecombinationFactor, o.PopulationSize)

		//Fillup
		population.Fillup(o.PopulationSize, o.Dimensions)

		//Evaluate using Approximator Function
		population.Evaluate(Approximator)

		//Sort
		population.Sort()

		//Append best to Hisory
		progressApproximated = append(progressApproximated, population[0].Fitness)

		cycle++
	}
	return
}
