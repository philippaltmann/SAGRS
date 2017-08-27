package Individual

import (
	"math"
	"math/rand"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Schaffer"
)

//TODO read from config file
const lowerBound float64 = Schaffer.LowerBound
const upperBound float64 = Schaffer.UpperBound

type Individual struct {
	Fitness float64
	Value   []float64
}

func GenerateIndividual(value []float64) (individual Individual) {
	return Individual{-1.0, value}
}

func GenerateRandomIndiviudal(dim int) (individual Individual) {
	var val []float64
	size := math.Abs(upperBound - lowerBound)
	for j := 0; j < dim; j++ {
		val = append(val, rand.Float64()*size+lowerBound)
	}
	return GenerateIndividual(val)
}

//func (individual *Individual) EvaluateFitness(fitnessFunction func([]float64) float64) {
func (individual *Individual) EvaluateFitness() {
	//TODO read evaluation function from config file
	individual.Fitness = Bohachevsky.EvaluateFitness(individual.Value)
}

func (individual Individual) Mutate() {

	position := rand.Intn(len(individual.Value))
	mutation := rand.Float64()*2 - 1 //Mutation -1 +1
	individual.Value[position] += mutation

}

func (individual Individual) Recombine(with Individual) (newIndividual Individual) {
	var newValue []float64
	count := len(individual.Value)
	for i := 0; i < count; i++ {
		tmpVal := (individual.Value[i] + with.Value[i]) / 2
		newValue = append(newValue, tmpVal)
	}
	newIndividual = GenerateIndividual(newValue)
	return
}
