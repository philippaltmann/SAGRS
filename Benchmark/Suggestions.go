package Benchmark

/*
import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/buger/goterm"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Ackley"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
)

func TestSuggestions(compare, rates, start int) {
	options := o.Options{
		Cycles:               1000,
		SuggestToEvaluation:  1,
		EvaluationPoolSize:   100,
		PopulationSize:       100,
		Dimensions:           2,
		SelectionFactor:      0.6,
		MutationFactor:       0.3,
		RecombinationFactor:  0.4,
		WriteProgress:        false,
		Verbose:              true,
		ResetPool:            false,
		ConvergenceThreshold: 100}

	approximatorType := "RBF"
	optimizationType := "Ackley"
	fitnessFunction := Ackley.EvaluateFitness

	var Approximator Approximation.Approximator
	switch approximatorType {
	case "LSM":
		Approximator = new(Approximation.LSMApproximator)
		options.EvaluationRate = 2
	case "RBF":
		Approximator = new(Approximation.RBFApproximator)
		options.EvaluationRate = 1
	default:
		Approximator = Approximation.GenerateEvaluation(fitnessFunction)
	}

	var resetPath string
	switch options.ResetPool {
	case true:
		resetPath = "Reset"
	default:
		resetPath = "NoReset"
	}

	_ = os.Mkdir("EvaluationRateTest"+approximatorType+"", 0777)
	_ = os.Mkdir("EvaluationRateTest"+approximatorType+"/"+resetPath+"", 0777)
	_ = os.Mkdir("EvaluationRateTest"+approximatorType+"/"+resetPath+"/"+optimizationType, 0777)
	_ = os.Mkdir("EvaluationRateTest"+approximatorType+"/"+resetPath+"/"+optimizationType+"/"+strconv.Itoa(options.Dimensions)+"Dim", 0777)
	basePath := "EvaluationRateTest" + approximatorType + "/" + resetPath + "/" + optimizationType + "/" + strconv.Itoa(options.Dimensions) + "Dim"

	file, _ := os.Create(basePath + "/EvaluationRateTestCompare.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Rate", "Best Fitness", "Convergence"})

	var functionFile *os.File

	goterm.Clear()
	for r := 0; r < rates; r++ {
		options.EvaluationRate = int(math.Pow(2, float64(r+start))) //0
		functionFile, _ = os.Create(basePath + "/Functions.txt")
		defer functionFile.Close()

		for i := 0; i < compare; i++ {
			goterm.MoveCursor(7, 1)
			goterm.Printf("Testing Rate %d: %d\n%s", options.EvaluationRate, i, makeProgressBar(i, compare, 50))
			goterm.Flush() // Call it every time at the end of rendering
			//options.ProgressFileName = "EvaluationRateTest" + aType + "/EvaluationRateTest" + strconv.Itoa(options.EvaluationRate) + "_" + strconv.Itoa(i)
			bestIndividual, cycle := a.Optimize(options, fitnessFunction, Approximator)
			line := []string{"Rate " + strconv.Itoa(options.EvaluationRate), strconv.FormatFloat(bestIndividual.Fitness, 'E', -1, 64), strconv.Itoa(cycle)}
			writer.Write(line)
			functionFile.WriteString(Approximator.GetFunction() + "\n\n")
		}
	}
}

func makeProgressBar(i, from, length int) string {
	done := int(float64(i) / float64(from) * float64(length))
	return "[" + strings.Repeat("=", done) + strings.Repeat("-", length-done) + "]"
}*/
