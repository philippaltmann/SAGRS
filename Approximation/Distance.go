package Approximation

import (
	"math"
	"sort"

	p "github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Population"
)

func ApproximateDistanceFitness(value []float64, ApproximationPool p.Population) (fitness float64) {
	const useFromCluster = 4
	var dc DistanceCluster
	var approximatedFitness float64 = 0
	var dsum float64 = 0.0

	for i := 0; i < len(ApproximationPool); i++ {
		dist := DistanceTo(value, ApproximationPool[i].Value)
		dc = append(dc, DistanceClusterEntry{
			Fitness:  ApproximationPool[i].Fitness,
			Distance: dist})
		if dist == math.Inf(+1) {
			//Matched individual: Set fitenss to max flaot to
			//prevent duplicates and higher diversity in evaluation pool
			return math.MaxFloat64
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
	/*if math.IsNaN(approximatedFitness / dsum) {
		panic(fmt.Sprintf("NaN resulting from %f / %f", approximatedFitness, dsum))
	}*/
	//Norm Approximated Fitness
	approximatedFitness /= dsum
	//approximatedFitness /= float64(len(evaluationPool))
	return approximatedFitness
}

type DistanceCluster []DistanceClusterEntry

type DistanceClusterEntry struct {
	Fitness  float64
	Distance float64
}

// Fitness sorter
func (f DistanceCluster) Len() int           { return len(f) }
func (f DistanceCluster) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f DistanceCluster) Less(i, j int) bool { return f[i].Distance > f[j].Distance }

func DistanceTo(value []float64, compare []float64) (distance float64) {
	distance = 0
	for v := 0; v < len(value); v++ {
		vi1 := math.Abs(value[v])
		vi2 := math.Abs(compare[v])
		//d := math.Sqrt(math.Pow(vi1-vi2, 2))
		d := math.Pow(vi1-vi2, 2)
		//fmt.Printf("\t Value %d (%f | %f): %f\n", v, individual.Value[v], compare.Value[v], d)
		distance += d

	}
	//distance /= float64(len(individual.Value))
	distance = 1 / distance
	return
}
