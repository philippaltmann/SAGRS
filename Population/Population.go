package Population

import (
	"math"
	"math/rand"
	"sort"
	"time"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

//Population is a Slice of individuals with custom applyable functions
type Population []i.Individual

// Helper functions for sorting capabilities
func (population Population) Len() int { return len(population) }
func (population Population) Swap(i, j int) {
	population[i], population[j] = population[j], population[i]
}
func (population Population) Less(i, j int) bool {
	return population[i].Fitness < population[j].Fitness
}

/*InitRandomPopulation generates a new random population
using the passed in size dimensions and value bounds*/
func InitRandomPopulation(size, dimensions int, min, max float64) (population Population) {
	for g := 0; g < size; g++ {
		population = append(population, i.GenerateRandomIndiviudal(dimensions, min, max))
	}
	return
}

/*func (population Population) Add(value []float64) Population {
	newIndividual := i.GenerateIndividual(value)
	population = append(population, newIndividual)
	return population
}*/

/*Evaluate calculates & assigns fitness values to all individuals
using the passed in function*/
func (population *Population) Evaluate(f func([]float64) float64) {
	for i := 0; i < len(*population); i++ {
		(*population)[i].Fitness = f((*population)[i].Value)
	}
	sort.Sort(Population(*population))
}

// Select performs natual Selection on the given Population
func (population *Population) Select(factor float64) {
	*population = (*population)[:int(float64(len(*population))*factor)]
}

/*Mutate Mutates individuals in the popultion with the given propability*/
func (population *Population) Mutate(factor float64) {
	// Start at position 1 to prevent mutation of best individual
	for p := 1; p < len(*population); p++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if r.Float64() > 1-factor {
			(*population)[p].Mutate()
		}
	}
}

/*Recombine appends Recombined Individuals with the passed in propability
to the Population until the maximum population size is reached*/
func (population *Population) Recombine(factor float64, size int) {
	for p := 0; p < len(*population); p++ {
		if rand.Float64() > 1-factor {
			combineWith := rand.Intn(len(*population))
			newIndividual := (*population)[p].Recombine((*population)[combineWith])
			if len(*population) < size {
				*population = append(*population, newIndividual)
			}
		}
	}
}

/*Fillup appends new random individuals in the given range to the
population until the maximum popultion size is reached*/
func (population *Population) Fillup(size, dimensions int, min, max float64) {
	for len(*population) < size {
		newIndiviudal := i.GenerateRandomIndiviudal(dimensions, min, max)
		*population = append(*population, newIndiviudal)
	}
}

/*GetDiversity calculates the populations diversity using the normed Euclidean Distance
between the Indivuduals, dividing by the population size*/
func (population Population) GetDiversity() (diversity float64) {
	for i := 0; i < len(population); i++ {
		for j := i; j < len(population); j++ {
			diversity += math.Sqrt(population[i].EuclideanDistance(population[j]))
		}
	}
	diversity /= float64(len(population))
	return
}
