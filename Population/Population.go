package Population

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"sort"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

type Population []i.Individual

func InitRandomPopulation(size int, dimensions int) (population Population) {
	for g := 0; g < size; g++ {
		population = append(population, i.GenerateRandomIndiviudal(dimensions))
	}
	return
}

func (g Population) Evaluate(f func([]float64) float64) {
	for i := 0; i < len(g); i++ {
		g[i].Fitness = f(g[i].Value)
	}
}

func (population Population) PrintBest(cycle int) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Printf("Cycle %d, with best at %f (%v)\n", cycle, population[0].Fitness, population[0].Value)
}

func (population *Population) Mutate(factor float64) {
	// Start at position 1 to prevent mutation of best individual
	for p := 1; p < len(*population); p++ {
		if rand.Float64() > 1-factor {
			(*population)[p].Mutate()
		}
	}
}

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

func (population *Population) Fillup(size int, dimensions int) {
	for len(*population) < size {
		newIndiviudal := i.GenerateRandomIndiviudal(dimensions)
		*population = append(*population, newIndiviudal)
	}
}

func (population Population) Sort() {
	sort.Sort(Population(population))
}

// Fitness sorter
func (f Population) Len() int           { return len(f) }
func (f Population) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f Population) Less(i, j int) bool { return math.Abs(f[i].Fitness) < math.Abs(f[j].Fitness) }
