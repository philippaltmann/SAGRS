package RBF

import "math"

//RadialBasisFunction Calculatable Function, center and width dependent
type RadialBasisFunction struct {
	center []float64
	width  float64
}

//Calculate applies the given value to the Radial Basis Function
func (f RadialBasisFunction) Calculate(value []float64) float64 {
	return 1 - math.Exp(-(EuclideanDistance(value, f.center))/(2*math.Pow(f.width, 2)))
}

//EuclideanDistance function for distance measure used in bias calculation
func EuclideanDistance(from []float64, to []float64) (distance float64) {
	for i, v := range from {
		distance += math.Pow((v - to[i]), 2)
	}
	return
}
