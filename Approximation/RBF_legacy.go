package Approximation
/*
func randomSamples(source [][]float64, k int) (samples [][]float64) {
	for i := 0; i < k; i++ {
		r := rand.Intn(len(source))
		samples = append(samples, source[r])
		source = append(source[:r], source[r+1:]...)
	}
	return
}

func equal(o, n [][]float64) bool {
	for i, c := range o {
		for j, v := range c {
			if n[i][j] != v {
				return false
			}
		}
	}
	return true
}

func assign(samples, centers [][]float64) (cluster [][][]float64) {
	cluster = make([][][]float64, len(centers))
	for _, v := range samples {
		d := math.MaxFloat64
		di := -1
		for i, c := range centers {
			dist := EuclideanDistance(v, c)
			if dist < d {
				d = dist
				di = i
			}
		}
		cluster[di] = append(cluster[di], v)
	}
	return
}

func adapt(centers [][]float64, cluster [][][]float64) (newCenters [][]float64) {
	newCenters = make([][]float64, len(centers))
	for i := range centers {
		for d := range centers[i] {

			newCenters[i] = append(newCenters[i], 0)
			for _, point := range cluster[i] {
				newCenters[i][d] += point[d]
			}
			newCenters[i][d] /= float64(len(cluster[i]))
		}
	}
	return
}
func Cluster(input [][]float64, k int) (centers [][]float64) {
	//Pick k random centers
	oldCenters := randomSamples(input, k)
	centers = randomSamples(input, k)
	var cluster [][][]float64
	for !equal(oldCenters, centers) {
		oldCenters = append([][]float64(nil), centers...)
		cluster = assign(input, centers)
		centers = adapt(oldCenters, cluster)
	}
	return
}
*/
