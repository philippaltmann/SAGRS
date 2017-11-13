package Approximation

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

type Approximator interface {
	Update(Population.Population)
	Predict([]float64) float64
	GetFunction() string
}

/*func (LSM LSMApproximator) Update(population Population.Population) {

}

func (LSM LSMApproximator) Predict(value []float64) {

}*/

func (RBF *RBFApproximator) Update(population Population.Population) {
	var input [][]float64
	var output []float64
	for _, i := range population {
		input = append(input, i.Value)
		output = append(output, i.Fitness)
	}
	RBF.UpdateRBFApproximator(input, output)
}

func (LSM *LSMApproximator) Update(population Population.Population) {
	LSM.ApproximationMatrix = GetLSMApproximator(population)
}

func (LSM *LSMApproximator) Predict(value []float64) float64 {
	return ApproximateFitness(value, LSM.ApproximationMatrix)
}

func (LSM *LSMApproximator) GetFunction() string {
	return fmt.Sprintln(mat64.Formatted(&LSM.ApproximationMatrix))
}

func (E *Evaluation) Update(population Population.Population) {

}

func (E *Evaluation) Predict(value []float64) float64 {
	return E.FitnessFunction(value)
}

func (E *Evaluation) GetFunction() string {
	return fmt.Sprintln(E.FitnessFunction)
}

type Evaluation struct {
	FitnessFunction func([]float64) float64
}

func GenerateEvaluation(fitnessFunction func([]float64) float64) *Evaluation {
	e := new(Evaluation)
	e.FitnessFunction = fitnessFunction
	return e
}
