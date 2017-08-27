package Approximation

import (
	"fmt"
	"math"
	"testing"

	i "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Individual"
)

func TestDistanceApproximation(t *testing.T) {
	testDimensions := 2
	testEvaluationPoolSize := 10
	var testEvaluationPool []i.Individual
	for t := 0; t < testEvaluationPoolSize; t++ {
		testEvaluationPool = append(
			testEvaluationPool, i.GenerateRandomIndiviudal(testDimensions))
		//testEvaluationPool[i].EvaluateFitness(FitnessFunctionTest)
		testEvaluationPool[t].EvaluateFitness()
		fmt.Printf("Pool Indivual %d: %v\n", t, testEvaluationPool[t])
	}

	ApproximationTestIndividual := i.GenerateRandomIndiviudal(testDimensions)
	fmt.Printf("Approximating Individual...\n")
	ApproximateDistanceFitness(ApproximationTestIndividual.Value, testEvaluationPool)
	fmt.Printf("Indivual: %v\n", ApproximationTestIndividual)

	//Test max value for duplicates
	testResul := ApproximateDistanceFitness(testEvaluationPool[0].Value, testEvaluationPool)

	if testResul != math.MaxFloat64 {
		t.Error("Not reacted to duplicates")
		t.Error(testResul)
		t.Fail()
	}
}
