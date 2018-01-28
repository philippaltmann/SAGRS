package Population

import (
	"math"
	"testing"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

var testPopulation Population

const populationSize int = 100
const dimensions int = 8
const min float64 = -100.0
const max float64 = 100.0

func testEvaluation(value []float64) (fitness float64) {
	fitness = 0.0
	for _, v := range value {
		fitness += math.Abs(v)
	}
	return
}

func TestRandomInit(t *testing.T) {
	testPopulation = InitRandomPopulation(populationSize, dimensions, min, max)
	if len(testPopulation) != populationSize {
		t.Error("Missmatching Population Size")
		t.Fail()
	}
	for i := 0; i < populationSize; i++ {
		if len(testPopulation[i].Value) != dimensions {
			t.Error("Missmatchting Dimensions")
			t.Fail()
		}
		for d := 0; d < dimensions; d++ {
			value := testPopulation[i].Value[d]
			if value > max || value < min {
				t.Error("Value not in Range")
				t.Fail()
			}
		}
	}
}

func TestEvaluateFitness(t *testing.T) {
	testPopulation.Evaluate(testEvaluation)
	for i := 0; i < populationSize; i++ {
		if testPopulation[i].Fitness != testEvaluation(testPopulation[i].Value) {
			t.Error("Evaluation Error")
			t.Fail()
		}
	}
}

func TestSort(t *testing.T) {
	//testPopulation.Sort()
	tmpFitness := -1.0
	for _, i := range testPopulation {
		if i.Fitness < tmpFitness {
			t.Error("Failed to Sort")
			t.Fail()
		}
		tmpFitness = i.Fitness
	}
}

func TestSelect(t *testing.T) {
	t.Log(testPopulation)
	testPopulation.Select(0.8)
	t.Log(testPopulation)
}

func TestDiversity(t *testing.T) {
	diversity := testPopulation.GetDiversity()
	if diversity < 0 {
		t.Error("Diversity not in Range")
		t.Log(diversity)
		t.Fail()
	}
}

func copyPopulation(population Population) Population {
	var populationCopy Population
	for _, individual := range population {
		populationCopy = append(populationCopy, i.DuplicateIndividual(individual))
	}
	return populationCopy
}

func TestMutation(t *testing.T) {
	testPopulationCopy := copyPopulation(testPopulation)
	testPopulation.Mutate(0)
	if testPopulation.GetDiversity() != testPopulationCopy.GetDiversity() {
		t.Error("Non working muation factor")
		t.Fail()
	}

	testPopulation.Mutate(1)
	if testPopulation.GetDiversity() == testPopulationCopy.GetDiversity() {
		t.Error("Did not mutate")
		t.Fail()
	}
}

func TestRecombination(t *testing.T) {
	testPopulation = testPopulation[:int(populationSize/2)]
	testPopulationCopy := copyPopulation(testPopulation)
	testPopulation.Recombine(0, populationSize)
	if len(testPopulation) != len(testPopulationCopy) {
		t.Error("Non working recombination factor")
		t.Fail()
	}

	testPopulation.Recombine(1, populationSize)
	if len(testPopulation) == len(testPopulationCopy) {
		t.Error("Did not recombine")
		t.Fail()
	}
}

func TestFillup(t *testing.T) {
	testPopulation = testPopulation[:int(populationSize/2)]
	testPopulation.Fillup(populationSize, dimensions, min, max)
	if len(testPopulation) != populationSize {
		t.Error("Fillup Failed")
		t.Fail()
	}
}
