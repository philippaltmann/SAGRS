package ApproximationOptimizer

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"

	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

//var ApproximationMatrix *mat64.Dense
var ApproximationMatrix a.RBFNet

//var ApproximationMatrix mat64.Dense

func Approximator(value []float64) float64 {
	//predictionInputMatrix := mat64.NewDense(1, len(value), value)
	//return n.Predict(ApproximationMatrix, predictionInputMatrix)
	return ApproximationMatrix.Predict(value)
	//return a.ApproximateRBF(ApproximationMatrix, value)
	//return a.ApproximateFitness(value, ApproximationMatrix)
}

func Optimize(o o.Options, FitnessFunction func([]float64) float64) (progressEvaluated []float64, progressApproximated []float64) {
	cycle := 0

	//Init, Evaluate & Sort Evaluation Pool random
	evaluationPool := g.InitRandomPopulation(o.EvaluationPoolSize, o.Dimensions)
	evaluationPool.Evaluate(FitnessFunction)
	evaluationPool.Sort()

	//Calculate Approximation Matrix for LSM Approximation
	//ApproximationMatrix = n.GetApproximator(evaluationPool)
	//ApproximationMatrix = a.GetLSMApproximator(evaluationPool)
	ApproximationMatrix = a.GetRBFApproximator(evaluationPool)

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
			//evaluationPool = evaluationPool[:len(evaluationPool)-1]

			//Update Approximator
			//ApproximationMatrix = n.GetApproximator(evaluationPool)
			//return
			ApproximationMatrix = a.GetRBFApproximator(evaluationPool)
			//ApproximationMatrix = a.GetLSMApproximator(evaluationPool)

			/*var distance float64
			for i := 0; i < len(evaluationPool); i++ {
				for j := i; j < len(evaluationPool); j++ {
					distance += a.DistanceTo(evaluationPool[i].Value, evaluationPool[j].Value)
				}
			}
			distance /= float64(len(evaluationPool))
			distance /= float64(len(evaluationPool))*/

			//Print Best
			fmt.Printf("Best approximated at %f, evaluated at %f\n\tBest in pool at %f, Size of pool: %d\n", population[0].Fitness, best.Fitness, evaluationPool[0].Fitness, len(evaluationPool))
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

		/*if population[0].Fitness == 0.0 {
			fmt.Println(population[0])
			//fmt.Println(mat64.Formatted(&ApproximationMatrix))
			fmt.Println(cycle)
			fmt.Println(population[len(population)-1])
			return
		}*/

		cycle++
	}
	evaluationPool.Sort()
	fmt.Print(ApproximationMatrix)
	fmt.Print("\n")
	fmt.Print(evaluationPool[0])
	fmt.Print("\n")
	fmt.Print(Approximator(evaluationPool[0].Value))
	fmt.Print("\n")

	//fmt.Println(mat64.Formatted(&ApproximationMatrix))
	//fmt.Println(ApproximationMatrix)

	file, _ := os.Create("quadraticTest3.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	cx := -20.0
	writer.Write([]string{"X", "Y", "Z"})
	for cx <= 20 {
		cy := -20.0

		for cy <= 20 {
			line := []string{toString(cx), toString(cy), toString(ApproximationMatrix.Predict([]float64{cx, cy}))}
			writer.Write(line)
			cy += 0.1
		}
		cx += 0.1
	}

	var distanceSum float64
	for i := 0; i < len(evaluationPool); i++ {
		for j := i; j < len(evaluationPool); j++ {
			distanceSum += a.EuclideanDistance(evaluationPool[i].Value, evaluationPool[j].Value)
		}
	}
	distanceSum = math.Sqrt(distanceSum)
	fmt.Println(distanceSum)
	distanceSum /= float64(len(evaluationPool))
	fmt.Println(distanceSum)

	return
}

func toString(num float64) string {
	return strconv.FormatFloat(num, 'f', 6, 64)
}
