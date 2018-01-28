package Individual

import "testing"

var lowerBound = -100.0
var upperBound = 100.0

func TestRest(t *testing.T) {
	dimensions := 32

	testIndidualOne := GenerateRandomIndiviudal(dimensions, lowerBound, upperBound)
	testIndidualTwo := GenerateRandomIndiviudal(dimensions, lowerBound, upperBound)
	testIndidualMut := DuplicateIndividual(testIndidualOne)

	for testIndidualOne.EuclideanDistance(testIndidualTwo) == 0 {
		testIndidualTwo = GenerateRandomIndiviudal(dimensions, lowerBound, upperBound)
	}

	//Test Distane
	distance := testIndidualOne.EuclideanDistance(testIndidualMut)
	if distance != 0 {
		t.Error("Non zero distance between identical Individuals")
		t.Fail()
	}

	//Test Mutation
	testIndidualMut.Mutate()
	distanceMut := testIndidualOne.EuclideanDistance(testIndidualMut)
	if distanceMut == 0 {
		t.Error("Mutation Failed")
		t.Fail()
	}

	//Test Recombination
	testIndidualRec := testIndidualOne.Recombine(testIndidualTwo)
	if len(testIndidualRec.Value) != dimensions {
		t.Error("Missmatching Dimensions")
		t.Fail()
	}
	distanceRec1 := testIndidualRec.EuclideanDistance(testIndidualOne)
	distanceRec2 := testIndidualRec.EuclideanDistance(testIndidualTwo)
	if distanceRec1 == 0 || distanceRec2 == 0 {
		t.Error("Recombination Failed")
		t.Fail()
	}
}

func TestRandomInit(t *testing.T) {
	testings := 100

	for r := 0; r < testings; r++ {
		testDimensions := 32
		testIndiviudal := GenerateRandomIndiviudal(testDimensions, lowerBound, upperBound)
		for d := 0; d < testDimensions; d++ {
			if testIndiviudal.Value[d] > upperBound || testIndiviudal.Value[d] < lowerBound {
				t.Errorf("Value not in range: %f | %d", testIndiviudal.Value[d], d)
			}
		}
	}
}
