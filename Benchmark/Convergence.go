package Benchmark

/*import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/buger/goterm"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Objective/Schwefel"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
)

func TestConvergence() {
	options := o.Options{
		Cycles:              20000,
		EvaluationRate:      2,
		SuggestToEvaluation: 1,
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          2,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4,
		WriteProgress:       false,
		Verbose:             true}

	//TODO track convergence if best fitness not changing
	//TODO adapt theshold to Suggested size
	approximatorType := ""
	optimizationType := "Schwefel"
	fitnessFunction := Schwefel.EvaluateFitness

	var Approximator Approximation.Approximator
	switch approximatorType {
	case "LSM":
		Approximator = new(Approximation.LSMApproximator)
		options.EvaluationRate = 2
	case "RBF":
		Approximator = new(Approximation.RBFApproximator)
		options.EvaluationRate = 8
	default:
		Approximator = Approximation.GenerateEvaluation(fitnessFunction)
	}

	_ = os.Mkdir("ConvergenceTest"+approximatorType+"", 0777)
	basePath := "ConvergenceTest" + approximatorType + "/" + optimizationType + "/" + strconv.Itoa(options.Dimensions) + "Dim"
	_ = os.Mkdir("ConvergenceTest"+approximatorType+"/"+optimizationType, 0777)
	_ = os.Mkdir("ConvergenceTest"+approximatorType+"/"+optimizationType+"/"+strconv.Itoa(options.Dimensions)+"Dim", 0777)

	file, _ := os.Create(basePath + "/SuggestedConvergingCompare.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Suggested", "Best Fitness", "Converged in cycle"})
	goterm.Clear()

	compare := 100
	for i := 1; i <= 64; i *= 2 {
		options.SuggestToEvaluation = i
		for j := 0; j < compare; j++ {
			goterm.MoveCursor(6, 1)
			goterm.Printf("Testing %d Sugestions\n%s\n", i, makeProgressBar(j, compare, 20))
			goterm.Flush()

			options.ProgressFileName = basePath + "/" + approximatorType + "TestRate" + strconv.Itoa(options.EvaluationRate)
			bestIndividual, cycle := a.Optimize(options, fitnessFunction, Approximator)
			writer.Write([]string{"Suggested " + strconv.Itoa(options.SuggestToEvaluation), strconv.FormatFloat(bestIndividual.Fitness, 'E', -1, 64), strconv.Itoa(cycle)})

		}
	}
}*/
