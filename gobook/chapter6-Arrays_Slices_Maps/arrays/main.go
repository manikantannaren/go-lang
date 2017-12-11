package main

import "fmt"

func main(){

    //Arrays
    var x [5]int
     x[4] = 100
     fmt.Println(x)

     //array length is of type int where as total is float.
     //use type conversoin
    var y [5]float64
    y[0] = 98
    y[1] = 93
    y[2] = 77
    y[3] = 82
    y[4] = 83
    var total float64 = 0
    for i := 0; i < 5; i++ {
       total += y[i]
    }
    fmt.Println(total / float64(len(y)))

    // using special for loop
    // the _ character tells compiler that the iteration variable is not used.
    // instead of _ if we use say i, then the compiler warns that i is not
    // being used
    // range can also be used with Strings to parse the characters
    // named range index can also be used instead of _.
    // Only catch the index must be used in subsequent statements
    for _, value:= range y{
        total += value
    }
    fmt.Println(total / float64(len(y)))
}
