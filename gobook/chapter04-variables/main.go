package main

import "fmt"

func main(){


  //parts of a variable declaration. type is optional as the compiler will warn
  var x string = "Hello world"
  fmt.Println(x)

  //declare and assign
  var y string
  y = x + " here"
  fmt.Println(y)
  y += " there"
  fmt.Println(y)

  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f",&input)
  output := input *2
  fmt.Print(output)
}
