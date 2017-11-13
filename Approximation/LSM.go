package Approximation

import (
	"math"

	"github.com/gonum/matrix/mat64"
	p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

type LSMApproximator struct {
	ApproximationMatrix mat64.Dense
}

func generateFitnessVector(population p.Population, avrg float64) *mat64.Dense {
	var fitnessSlice []float64
	for _, individual := range population {
		fitnessSlice = append(fitnessSlice, individual.Fitness-avrg)
	}
	fitnessVector := mat64.NewDense(len(population), 1, fitnessSlice)
	return fitnessVector
}

func generateValueMatrix(population p.Population, avrg []float64) *mat64.Dense {
	var valueSlice []float64
	test := 0.0
	for _, individual := range population {
		for i, value := range individual.Value {
			test += value - avrg[i]
			valueSlice = append(valueSlice, value-avrg[i])
		}
	}
	valueMatrix := mat64.NewDense(len(population), len(valueSlice)/len(population), valueSlice)
	return valueMatrix
}

func generateAverages(population p.Population) (avrgX []float64, avrgY float64) {
	dimensions := len(population[0].Value)
	avrgY = 0

	for i := 0; i < dimensions; i++ {
		avrgX = append(avrgX, 0)
		for _, individual := range population {
			if i == 0 {
				//Sum avrgY
				avrgY += individual.Fitness
			}
			avrgX[i] += individual.Value[i]
		}
		avrgX[i] /= float64(len(population))

	}
	avrgY /= float64(len(population))
	return
}

func GetLSMApproximator(population p.Population) (approxmationMatrix mat64.Dense) {

	avrgX, avrgY := generateAverages(population)
	y := generateFitnessVector(population, avrgY)
	X := generateValueMatrix(population, avrgX)
	XT := X.T()
	var first mat64.Dense // construct a new zero-sized matrix
	first.Mul(XT, X)

	//r, c := first.Dims()
	var firstInverse mat64.Dense
	firstInverse.Inverse(&first)

	var second mat64.Dense
	second.Mul(XT, y)

	var theta mat64.Dense
	theta.Mul(&firstInverse, &second)

	/*var aSlice []float64
	for i := 0; i < len(population[0].Value); i++ {
		a := avrgY - theta.At(i, 0)*avrgX[i]
		aSlice = append(aSlice, a)
	}*/

	//r, c := theta.Dims()
	//fmt.Printf("R: %d, C: %d\n", r, c)
	/*tmpMatrix := theta.Grow(0, 1)
	thetaMulti := *mat64.DenseCopyOf(tmpMatrix)*/
	//r, c = thetaMulti.Dims()
	//fmt.Printf("R: %d, C: %d\n", r, c)

	//fmt.Print(thetaMulti.At(0, 1))
	//thetaMulti.SetCol(1, aSlice)
	//fmt.Print(thetaMulti.At(0, 1))

	//fmt.Printf("TM: %v\n", thetaMulti)
	return theta

}

func ApproximateFitness(value []float64, ApproximationMatrix mat64.Dense) (fitness float64) {
	valueVector := mat64.NewDense(1, len(value), value)
	var resultMatrix mat64.Dense
	resultMatrix.Mul(valueVector, &ApproximationMatrix)
	result := math.Abs(resultMatrix.At(0, 0))

	//result := 0.0 //resultMatrix.At(0, 0)
	/*deg := 1
	for d := 0; d < 2; d++ {
		for i := 0; i < len(value); i++ {
			result += math.Pow(value[i], float64(deg)) * ApproximationMatrix.At(i, d)
		}
		deg--
	}*/
	return result
}
