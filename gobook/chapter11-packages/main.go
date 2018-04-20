package main

import (
	"fmt"
	"gobook/chapter11-packages/mymath"
)

func main() {
xs:
	[]float64{1, 2, 3, 4}
	avg := mymath.average(xs)
	fmt.Println(avg)
}
