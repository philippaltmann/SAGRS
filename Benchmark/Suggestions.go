package Benchmark

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/buger/goterm"
	e "github.com/philipp-altmann/SAGRS/Environment"
	r "github.com/philipp-altmann/SAGRS/Recommender"
)

//TestSuggestions Runs tests for different amounts of suggeted individuals
func TestSuggestions(compare, rate int, reset bool, suggestions []int, approximator string, objective string, cycles int) {
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
		Approximator:        approximator,
		Objective:           objective}

	path := "Tests/Suggestions"
	_ = os.Mkdir(path, 0777)
	path += "/" + environment.Approximator
	_ = os.Mkdir(path, 0777)
	path += "/" + environment.Objective
	_ = os.Mkdir(path, 0777)

	environment.Dump(path + "/Environment.json")

	file, _ := os.Create(path + "/SuggestionsTestCompare.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"Suggestions", "Best Fitness", "Convergence"})
	writer.Flush()

	goterm.Clear()
	for _, s := range suggestions {
		environment.SuggestToEvaluation = s

		for i := 0; i < compare; i++ {
			environment.ProgressFileName = path + "/Progress_" + strconv.Itoa(environment.SuggestToEvaluation) + "_" + strconv.Itoa(i)
			goterm.MoveCursor(7, 1)
			goterm.Printf("Testing %d Suggestions: %d\n%s", environment.SuggestToEvaluation, i, makeProgressBar(i, compare, 50))
			goterm.Flush()
			bestIndividual, cycle := r.Run(environment)
			line := []string{strconv.Itoa(environment.SuggestToEvaluation), strconv.FormatFloat(bestIndividual.Fitness, 'E', -1, 64), strconv.Itoa(cycle)}
			writer.Write(line)
			writer.Flush()
		}
	}
}
