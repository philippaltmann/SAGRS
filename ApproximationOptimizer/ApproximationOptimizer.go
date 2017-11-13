package ApproximationOptimizer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/buger/goterm"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	g "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func Optimize(o o.Options, FitnessFunction func([]float64) float64, Approximator a.Approximator /*, initialPopulation g.Population*/) ( /*bestPool g.Population */ bestIndividual i.Individual, cycle int) {

	//Init, Evaluate & Sort Evaluation Pool random
	evaluationPool := g.InitRandomPopulation(o.EvaluationPoolSize, o.Dimensions)
	evaluationPool.Evaluate(FitnessFunction)
	evaluationPool.Sort()

	convergenceCycle := o.Cycles

	/*if len(initialPopulation) == o.EvaluationPoolSize {
		evaluationPool = initialPopulation
	}*/
	functionFile, _ := os.Create("FunctionsProgress.txt")
	defer functionFile.Close()

	//Update Approximator
	Approximator.Update(evaluationPool)

	var writer *csv.Writer
	if o.WriteProgress {
		file, _ := os.Create(o.ProgressFileName + ".csv")
		defer file.Close()
		writer = csv.NewWriter(file)
		defer writer.Flush()
		writer.Write([]string{"Cycle", "Best in Pool", "Approximation Success", "Approximation Error"})
	}

	//Init, Approximate & Sort Population
	population := g.InitRandomPopulation(o.PopulationSize, o.Dimensions)
	population.Evaluate(Approximator.Predict)
	population.Sort()

	converging := 0 //0 for not number for unsuccessful cycles in a row
	for cycle := 0; cycle < o.Cycles; cycle++ {
		if o.ResetPool {
			//Init, Approximate & Sort Population
			population := g.InitRandomPopulation(o.PopulationSize, o.Dimensions)
			population.Evaluate(Approximator.Predict)
			population.Sort()
		}

		//Optimize Population for Suggestion using Approximator
		for generation := 0; generation < o.EvaluationRate; generation++ {

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
			population.Evaluate(Approximator.Predict)

			//Sort
			population.Sort()
		}

		//Suggest
		var suggestions []i.Individual
		bestApproximated := population[0].Fitness
		better := 0
		worse := 0
		for suggestion := 0; suggestion < o.SuggestToEvaluation; suggestion++ {
			individual := i.GenerateIndividual(population[suggestion].Value)
			individual.Fitness = FitnessFunction(individual.Value)
			suggestions = append(suggestions, individual)
			if individual.Fitness > evaluationPool[len(evaluationPool)-1].Fitness {
				worse++
			}
			if individual.Fitness < evaluationPool[0].Fitness {
				better++
			}
		}
		evaluationPool = append(evaluationPool, suggestions...)

		evaluationPool.Evaluate(FitnessFunction)
		evaluationPool.Sort()

		//Cut worst to keep Evaluation Pool Size Constant
		evaluationPool = evaluationPool[:len(evaluationPool)-o.SuggestToEvaluation]

		//Update Approximator
		Approximator.Update(evaluationPool)

		if o.Verbose {
			//Print Stats
			goterm.MoveCursor(1, 1)
			betterString := goterm.Color(fmt.Sprintf("[%s]Better", getBetterLamp(better)), goterm.CYAN)
			goterm.Printf("%s[%4d]\n", makeProgressBar(cycle, o.Cycles, 50, goterm.GREEN), cycle)
			goterm.Printf("%s   Suggestion Success: %s\n", betterString, makeProgressBar(o.SuggestToEvaluation-worse, o.SuggestToEvaluation, 20, goterm.BLUE))
			goterm.Printf("Best in Pool at %6.6f       |Approximated at %6.6f\n", evaluationPool[0].Fitness, Approximator.Predict(evaluationPool[0].Value))

			goterm.Printf("Rejected %d of %d\n", worse, o.SuggestToEvaluation)
			goterm.Flush() // Call it every time at the end of rendering
		}

		terminateWhenConverging := true
		functionFile.WriteString(Approximator.GetFunction() + "\n\n")

		if terminateWhenConverging {
			if worse == o.SuggestToEvaluation || evaluationPool[0].Fitness == 0.0 {
				if converging > 0 {
					converging++
					/*goterm.MoveCursor(4, 1)
					goterm.Printf("Convergin for %d Cycles", converging)
					goterm.Flush()*/
				} else {
					converging = 1
					convergenceCycle = cycle
				}
			} else {
				converging = 0
				convergenceCycle = cycle
			}
		}
		if converging == o.ConvergenceThreshold {
			//return evaluationPool[0], cycle //evaluationPool
		}

		//goterm.MoveCursor(3, 1)
		//goterm.Println(makeProgressBar(cycle, o.Cycles, 50))

		//Select
		/*population = population[o.SuggestToEvaluation:]
		fmt.Println(len(population))

		//Fillup
		population.Fillup(o.PopulationSize, o.Dimensions)*/

		if o.WriteProgress {
			//Write Progress to csv
			line := make([]string, 4)
			//Cycle
			line[0] = strconv.Itoa(cycle)

			//Best in Pool
			line[1] = strconv.FormatFloat(evaluationPool[0].Fitness, 'E', -1, 64)

			//Approximation Success
			line[2] = strconv.FormatFloat((float64(o.SuggestToEvaluation-worse) / float64(o.SuggestToEvaluation)), 'E', -1, 64)

			//Approximation Error
			line[3] = strconv.FormatFloat((bestApproximated - FitnessFunction(population[0].Value)), 'E', -1, 64)

			writer.Write(line)
		}
	}
	return evaluationPool[0], convergenceCycle //evaluationPool
}

func getBetterLamp(better int) string {
	if better > 0 {
		return "•"
	}
	return " "
}

func makeProgressBar(i, from, length int, color int) string {
	done := int(float64(i) / float64(from) * float64(length))
	return "[" + goterm.Color(strings.Repeat("=", done), color) + strings.Repeat("-", length-done) + "]"
}
