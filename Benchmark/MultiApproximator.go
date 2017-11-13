package Benchmark

/*import (
	"os"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Approximation"
	a "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/ApproximationOptimizer"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
	o "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Options"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func MultiApproximate() {
	options := o.Options{
		Cycles:              20000,
		EvaluationRate:      2,
		SuggestToEvaluation: 1,
		EvaluationPoolSize:  100,
		PopulationSize:      100,
		Dimensions:          32,
		SelectionFactor:     0.6,
		MutationFactor:      0.3,
		RecombinationFactor: 0.4,
		WriteProgress:       true,
		Verbose:             true}

	//TODO track convergence if best fitness not changing
	//TODO adapt theshold to Suggested size
	_ = os.Mkdir("MultiApproximatorTest", 0777)
	fitnessFunction := Bohachevsky.EvaluateFitness

	Approximator1 := new(Approximation.LSMApproximator)
	options.EvaluationRate = 2

	initialPopulation := *new(Population.Population)

	options.ProgressFileName = "MultiApproximatorTest/LSM"
	bestPool, _ := a.Optimize(options, fitnessFunction, Approximator1, initialPopulation)

	Approximator2 := new(Approximation.RBFApproximator)
	options.EvaluationRate = 8
	options.ProgressFileName = "MultiApproximatorTest/RBF"
	a.Optimize(options, fitnessFunction, Approximator2, bestPool)

}*/
