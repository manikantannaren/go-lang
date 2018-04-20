package mymath

func average(xs []float64) (retValue float64) {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	retValue = total / float64(len(xs))
	return
}
