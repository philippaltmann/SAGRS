package Generation

import (
	"fmt"
	"os"
	"os/exec"
	"sort"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

type Generation []i.Individual

func (generation Generation) PrintBest(cycle int) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Printf("Cycle %d, with best at %f (%v)\n", cycle, generation[0].Fitness, generation[0].Value)
}

func (generation Generation) Sort() {
	sort.Sort(Generation(generation))
}

// Fitness sorter
func (f Generation) Len() int           { return len(f) }
func (f Generation) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f Generation) Less(i, j int) bool { return f[i].Fitness < f[j].Fitness }
