package Benchmark

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"

	"github.com/buger/goterm"
	e "github.com/philipp-altmann/SAGRS/Environment"
	g "github.com/philipp-altmann/SAGRS/GeneticOptimizer"
	r "github.com/philipp-altmann/SAGRS/Recommender"
)

//Compare Benchmark SAGRS against Genetic Algorithm and Random Search
func Compare(compare, rate int, reset bool, suggestions int, approximator string, objective string, cycles int) {

	environment := e.Environment{
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          2,
		SelectionFactor:     0.9,
		MutationFactor:      0.1,
		RecombinationFactor: 0.05,
		WriteProgress:       true,
		Verbose:             true,
		ResetPool:           reset,
		EvaluationRate:      rate,
		Cycles:              cycles,
		SuggestToEvaluation: suggestions,
		Approximator:        approximator,
		Objective:           objective}

	randomEnv := e.Environment{
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          2,
		SelectionFactor:     0.9,
		MutationFactor:      0.1,
		RecombinationFactor: 0.05,
		WriteProgress:       true,
		Verbose:             true,
		EvaluationRate:      0,
		ResetPool:           true,
		Cycles:              cycles,
		SuggestToEvaluation: suggestions,
		Approximator:        approximator,
		Objective:           objective}

	realEvaluations := environment.EvaluationPoolSize + cycles*suggestions
	geneticEnv := e.Environment{
		PopulationSize:      int(math.Sqrt(float64(realEvaluations))),
		Cycles:              int(math.Sqrt(float64(realEvaluations))),
		Dimensions:          environment.Dimensions,
		SelectionFactor:     environment.SelectionFactor,
		MutationFactor:      environment.MutationFactor,
		RecombinationFactor: environment.RecombinationFactor,
		WriteProgress:       true,
		Verbose:             true,
		EvaluationRate:      rate,
		Approximator:        approximator,
		Objective:           objective}

	path := "Tests/Comparison"
	_ = os.Mkdir(path, 0777)
	path += "/" + environment.Approximator
	_ = os.Mkdir(path, 0777)
	path += "/" + environment.Objective
	_ = os.Mkdir(path, 0777)

	environment.Dump(path + "/Environment.json")

	file, _ := os.Create(path + "/Compare.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"Best Fitness SAGRS", "Best Fitness GA", "Best Fitness RR"})
	writer.Flush()
	goterm.Clear()

	for i := 0; i < compare; i++ {
		goterm.MoveCursor(7, 1)
		goterm.Printf("Comparing SAGRS (%s | %s): %d\n%s", objective, approximator, i, makeProgressBar(i, compare, 50))
		goterm.Flush() // Call it every time at the end of rendering
		environment.ProgressFileName = path + "/Progress_SAGRS_" + strconv.Itoa(i)
		bestIndividualSAGRS, _ := r.Run(environment)

		goterm.MoveCursor(7, 1)
		goterm.Printf("Comparing GA (%s | %s): %d\n%s", objective, approximator, i, makeProgressBar(i, compare, 50))
		goterm.Flush() // Call it every time at the end of rendering
		geneticEnv.ProgressFileName = path + "/Progress_GA_" + strconv.Itoa(i)
		bestIndividualGA := g.Optimize(geneticEnv)

		goterm.MoveCursor(7, 1)
		goterm.Printf("Comparing RR (%s | %s): %d\n%s", objective, approximator, i, makeProgressBar(i, compare, 50))
		goterm.Flush() // Call it every time at the end of rendering
		randomEnv.ProgressFileName = path + "/Progress_RR_" + strconv.Itoa(i)
		bestIndividualRR, _ := r.Run(randomEnv)

		line := []string{strconv.FormatFloat(bestIndividualSAGRS.Fitness, 'E', -1, 64), strconv.FormatFloat(bestIndividualGA.Fitness, 'E', -1, 64), strconv.FormatFloat(bestIndividualRR.Fitness, 'E', -1, 64)}
		writer.Write(line)
		writer.Flush()
	}
}
