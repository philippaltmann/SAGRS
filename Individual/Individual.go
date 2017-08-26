package Individual

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Bohachevsky"
)

//TODO read from config file
const lowerBound float64 = Bohachevsky.LowerBound
const upperBound float64 = Bohachevsky.UpperBound

type Individual struct {
	Fitness float64
	Value   []float64
}

func GenerateIndividual(value []float64) (individual Individual) {
	return Individual{-1.0, value}
}

func GenerateRandomIndiviudal(dim int) (individual Individual) {
	var val []float64
	size := math.Abs(upperBound - lowerBound)
	for j := 0; j < dim; j++ {
		val = append(val, rand.Float64()*size+lowerBound)
	}
	return GenerateIndividual(val)
}

func (individual *Individual) EvaluateFitness() {
	//TODO read evaluation function from config file
	individual.Fitness = Bohachevsky.EvaluateFitness(individual.Value)
}

func (individual *Individual) ApproximateFitness(evaluationPool []Individual) {
	const useFromCluster = 4
	var dc DistanceCluster
	var approximatedFitness float64 = 0
	var dsum float64 = 0.0

	for i := 0; i < len(evaluationPool); i++ {
		//fmt.Printf("Approxmation %d:\n", i)
		dist := individual.DistanceTo(evaluationPool[i])
		dc = append(dc, DistanceClusterEntry{
			Fitness:  evaluationPool[i].Fitness,
			Distance: dist})
		if dist == math.Inf(+1) {
			//Matched individual: Set fitenss to max flaot to
			//prevent duplicates and higher diversity in evaluation pool
			approximatedFitness = math.MaxFloat64
			goto SetFitness
		}
	}
	sort.Sort(dc)

	for i := 0; i < useFromCluster; i++ {
		dsum += dc[i].Distance
		//fmt.Printf("\t-> %f * evaluationPool[i].Fitness || Increasing by %f\n", dist, dist*evaluationPool[i].Fitness)
		approximatedFitness += dc[i].Distance * dc[i].Fitness
	}
	//fmt.Printf("Sum of Distance: %f", dsum)
	//Norm fittness to be on same scale
	if math.IsNaN(approximatedFitness / dsum) {
		panic(fmt.Sprintf("NaN resulting from %f / %f", approximatedFitness, dsum))
	}
	//Norm Approximated Fitness
	approximatedFitness /= dsum
	//approximatedFitness /= float64(len(evaluationPool))
SetFitness:
	individual.Fitness = approximatedFitness
}

type DistanceCluster []DistanceClusterEntry

type DistanceClusterEntry struct {
	Fitness  float64
	Distance float64
}

func (dc DistanceCluster) SortDistanceCluster() {
	sort.Sort(DistanceCluster(dc))
}

// Fitness sorter
func (f DistanceCluster) Len() int           { return len(f) }
func (f DistanceCluster) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f DistanceCluster) Less(i, j int) bool { return f[i].Distance > f[j].Distance }

func (individual Individual) DistanceTo(compare Individual) (distance float64) {
	distance = 0
	for v := 0; v < len(individual.Value); v++ {
		vi1 := math.Abs(individual.Value[v])
		vi2 := math.Abs(compare.Value[v])
		//d := math.Sqrt(math.Pow(vi1-vi2, 2))
		d := math.Pow(vi1-vi2, 2)
		//fmt.Printf("\t Value %d (%f | %f): %f\n", v, individual.Value[v], compare.Value[v], d)
		distance += d

	}
	//distance /= float64(len(individual.Value))
	distance = 1 / distance
	return
}

func (individual Individual) Mutate() {
	position := rand.Intn(len(individual.Value))
	mutation := rand.Float64()*2 - 1 //Mutation -1 +1
	individual.Value[position] += mutation
}

func (individual Individual) Recombine(with Individual) (newIndividual Individual) {
	var newValue []float64
	count := len(individual.Value)
	for i := 0; i < count; i++ {
		tmpVal := (individual.Value[i] + with.Value[i]) / 2
		newValue = append(newValue, tmpVal)
	}
	newIndividual = GenerateIndividual(newValue)
	return
}
