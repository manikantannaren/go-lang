package main

import "fmt"

func main(){
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
}

//function name, parameter name and type....return type (with optional return type name)
func average(scores []float64) (avgscore float64) {
    total := 0.0
    for _, v := range scores {
        total += v
    }
    avgscore = total / float64(len(scores))
    return //not required to name the return variable here since already declared
    //this also works
    //return avgscore
}
