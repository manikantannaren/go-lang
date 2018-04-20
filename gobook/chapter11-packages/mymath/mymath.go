package mymath

//Average average of floating values
func Average(xs []float64) (retValue float64) {
	//name Average has to be init caps per Go conventions for packaged routines
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	retValue = total / float64(len(xs))
	return
}
