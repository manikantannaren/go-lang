package main

import (
	"fmt"
	//requires full path to package
	"github.com/manikantannaren/golang-learning/gobook/chapter11-packages/mymath"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := mymath.Average(xs)
	fmt.Println(avg)
}
