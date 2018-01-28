package Benchmark

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/buger/goterm"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	e "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Environment"
)

func TestEvaluationRates(compare int, rates []int, approximator string, objective string, cycles int) {
	environment := e.Environment{
		SuggestToEvaluation: 1,
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          2,
		SelectionFactor:     0.9,
		MutationFactor:      0.1,
		RecombinationFactor: 0.05,
		WriteProgress:       true,
		Verbose:             true,
		ResetPool:           false,
		Cycles:              cycles,
		Approximator:        approximator,
		Objective:           objective}

	var resetPath string
	switch environment.ResetPool {
	case true:
		resetPath = "Reset"
	default:
		resetPath = "NoReset"
	}
	path := "Tests/EvaluationRate"
	_ = os.Mkdir(path, 0777)
	path += "/" + strconv.Itoa(environment.Dimensions) + "Dim"
	_ = os.Mkdir(path, 0777)
	path += "/" + resetPath
	_ = os.Mkdir(path, 0777)
	path += "/" + environment.Approximator
	_ = os.Mkdir(path, 0777)
	path += "/" + environment.Objective
	_ = os.Mkdir(path, 0777)

	scan, _ := ioutil.ReadDir(path)
	dirs := make([]string, 0)
	for _, f := range scan {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}

	dir, _ := strconv.Atoi(dirs[len(dirs)-1])
	path += "/" + strconv.Itoa(1+dir)
	_ = os.Mkdir(path, 0777)

	environment.Dump(path + "/Environment.json")

	file, _ := os.Create(path + "/EvaluationRateTestCompare.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"Rate", "Best Fitness", "Convergence"})
	writer.Flush()

	goterm.Clear()
	for _, r := range rates {
		environment.EvaluationRate = r //int(math.Pow(2, float64(r+start))) //0

		for i := 0; i < compare; i++ {
			environment.ProgressFileName = path + "/Progress_" + strconv.Itoa(environment.EvaluationRate) + "_" + strconv.Itoa(i)
			goterm.MoveCursor(7, 1)
			goterm.Printf("Testing Rate %d: %d\n%s", environment.EvaluationRate, i, makeProgressBar(i, compare, 50))
			goterm.Flush() // Call it every time at the end of rendering
			//options.ProgressFileName = "EvaluationRateTest" + aType + "/EvaluationRateTest" + strconv.Itoa(options.EvaluationRate) + "_" + strconv.Itoa(i)
			bestIndividual, cycle := a.Optimize(environment)
			line := []string{"Rate " + strconv.Itoa(environment.EvaluationRate), strconv.FormatFloat(bestIndividual.Fitness, 'E', -1, 64), strconv.Itoa(cycle)}
			writer.Write(line)
			writer.Flush()

		}
	}
}

func makeProgressBar(i, from, length int) string {
	done := int(float64(i) / float64(from) * float64(length))
	return "[" + strings.Repeat("=", done) + strings.Repeat("-", length-done) + "]"
}
