package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	s1 := new(Rectangle)
	s1.length = 10.2
	s1.width = 3.4
	calculateAndDisplay(s1)
	s2 := new(Circle)
	s2.radius = 5.5
	calculateAndDisplay(s2)
}

//Shape defines generic methods
// Implementations just define the methods
type Shape interface {
	area() float64
	perimeter() float64
	dimensions() map[string]float64
}

// Rectangle defines a shape of Rectangle
// Above Comment required by golint
type Rectangle struct {
	Shape
	length, width float64
}

// Circle defines a shape of Circle
type Circle struct {
	Shape
	radius float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}
func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func (r *Rectangle) dimensions() (x map[string]float64) {
	x = make(map[string]float64)
	x["length"] = r.length
	x["width"] = r.width
	return
}
func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c *Circle) dimensions() (x map[string]float64) {
	x = make(map[string]float64)
	x["radius"] = c.radius
	return x
}

/*
 * Go lang does not advocate calling functions on pointers
 */

func dimensionAsString(s Shape) (retValue string) {
	var dimensions = s.dimensions()
	concatString := ""
	for k, v := range dimensions {
		concatString += k + ": " + strconv.FormatFloat(v, 'f', 2, 64) + ", "
	}
	retValue = strings.Trim(concatString, ", ")
	return
}

func calculateAndDisplay(s Shape) {
	areaOfShape := strconv.FormatFloat(s.area(), 'f', 2, 64)
	perimeterOfShape := strconv.FormatFloat(s.perimeter(), 'f', 2, 64)
	dim := dimensionAsString(s)

	fmt.Println("Shape {Dimensions:", dim, ", Area:", areaOfShape, ", Perimeter:", perimeterOfShape, "}")

}
