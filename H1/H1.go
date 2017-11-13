package H1

import "math"

const UpperBound float64 = 100.0
const LowerBound float64 = -100.0

func EvaluateFitness(values []float64) float64 {
	x1 := values[0]
	x2 := values[1]

	return -(math.Pow(math.Sin(x1-(x2/8)), 2)+math.Pow(math.Sin(x2+x1/8), 2))/(math.Sqrt(math.Pow((x1-8.6998), 2)+math.Pow((x2-6.7665), 2))+1) + 2
}
