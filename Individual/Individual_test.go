package Individual

import (
	"fmt"
	"testing"
)

func TestApproximation(t *testing.T) {
	testDimensions := 2
	testEvaluationPoolSize := 10
	var testEvaluationPool []Individual
	for i := 0; i < testEvaluationPoolSize; i++ {
		testEvaluationPool = append(
			testEvaluationPool, GenerateRandomIndiviudal(testDimensions))
		testEvaluationPool[i].EvaluateFitness()
		fmt.Printf("Pool Indivual %d: %v\n", i, testEvaluationPool[i])
	}

	ApproximationTestIndividual := GenerateRandomIndiviudal(testDimensions)
	fmt.Printf("Approximating Individual...\n")
	ApproximationTestIndividual.ApproximateFitness(testEvaluationPool)
	fmt.Printf("Indivual: %v\n", ApproximationTestIndividual)

}
