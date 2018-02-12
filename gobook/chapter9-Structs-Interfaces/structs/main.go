package main

import (
	"fmt"
	"math"
)

func main() {
	c1 := Circle{x: 0.0, y: 0.0, r: 3.0}
	areaOfCircle := calculateArea(&c1)
	fmt.Println("Area of circle with radius :", c1.r, areaOfCircle)

	c2 := c1
	areaOfCircle = c2.area()
	fmt.Println("Area computed using receiver method", areaOfCircle)

	h1 := new(Humanoid)
	h1.name = "Marvin"
	h1.version = "1.1"
	fmt.Println("Hi, my version is", h1.version)
	h1.Talk()
}
func calculateArea(circle *Circle) (retArea float64) {
	retArea = math.Pi * circle.r * circle.r
	return
}

//reciever method. method name comes after paramters
func (circle *Circle) area() (retArea float64) {
	retArea = math.Pi * circle.r * circle.r
	return
}

// Circle defines a shape of circle
// Above Comment required by golint
type Circle struct {
	x float64
	y float64
	r float64
}

type Person struct {
	name string
}

type Humanoid struct {
	Person  //Humanoid IS A type of Person (inheritance)
	version string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.name)

}
