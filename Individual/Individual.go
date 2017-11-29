package Individual

import (
	"math"
	"math/rand"
	"time"
)

//Individual Type
type Individual struct {
	Fitness float64
	Value   []float64
}

//GenerateIndividual with given values
func GenerateIndividual(value []float64) (individual Individual) {
	return Individual{-1.0, value}
}

//GenerateRandomIndiviudal with given dimensions
func GenerateRandomIndiviudal(dim int, min, max float64) Individual {
	var value []float64
	size := math.Abs(max - min)
	for j := 0; j < dim; j++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		value = append(value, r.Float64()*size+min)
	}
	return GenerateIndividual(value)
}

//Mutate the receiver Individual  by ±1 in a random dimension
func (individual *Individual) Mutate() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	position := r.Intn(len(individual.Value))
	mutation := r.Float64()*2 - 1 //Mutation -1 +1
	individual.Value[position] += mutation
}

//Recombine the receiver Individual with the passed in using crossover
func (individual Individual) Recombine(with Individual) (newIndividual Individual) {
	value1 := make([]float64, len(individual.Value))
	value2 := make([]float64, len(with.Value))
	var newValue []float64
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	crossover := r.Intn(len(individual.Value)-1) + 1
	copy(value1, individual.Value)
	copy(value2, with.Value)
	newValue = append(value1[:crossover], value2[crossover:]...)
	newIndividual = GenerateIndividual(newValue)
	return
}

//DuplicateIndividual Deep Copies Individuals
func DuplicateIndividual(individual Individual) Individual {
	valueCopy := make([]float64, len(individual.Value))
	copy(valueCopy, individual.Value)
	duplicate := GenerateIndividual(valueCopy)
	return duplicate
}

//EuclideanDistance calculated between the receiver and passed in Individual
func (individual Individual) EuclideanDistance(to Individual) (distance float64) {
	for i, v := range individual.Value {
		distance += math.Pow((v - to.Value[i]), 2)
	}
	return
}
