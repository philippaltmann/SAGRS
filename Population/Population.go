package Population

import (
	"fmt"
	"os"
	"os/exec"
	"sort"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

type Population []i.Individual

func InitRandomPopulation(size int, dimensions int) (population Population) {
	for g := 0; g < size; g++ {
		population = append(population, i.GenerateRandomIndiviudal(dimensions))
		population[g].EvaluateFitness()
	}
	return
}

/*func (g Population) EvaluateFitness(f func([]float64)) {
	for i := 0; i < len(g); i++ {
		g[i].Fitness = f(g[i].Value)
	}
}*/

func (population Population) PrintBest(cycle int) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Printf("Cycle %d, with best at %f (%v)\n", cycle, population[0].Fitness, population[0].Value)
}

func (population Population) Sort() {
	sort.Sort(Population(population))
}

// Fitness sorter
func (f Population) Len() int           { return len(f) }
func (f Population) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f Population) Less(i, j int) bool { return f[i].Fitness < f[j].Fitness }
