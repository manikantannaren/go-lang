package main

import "fmt"

func main(){
    x,y := returnMultipleValues()
    fmt.Println("Point x : ", x, "Point y: ",y)
    variadicParams(10,20,30,40)
    sum := addUsingVariadicParams(10,20,30,40)
    fmt.Println("Total using variadic ",sum)

    slicedValus := []int{5,15,25,35}
    sum = addUsingVariadicParams(slicedValus...)
    fmt.Println("Total using slice and variadic ",sum)
}

func returnMultipleValues () (pointx int, pointy int){
    return 5,6
}

func variadicParams(values ...int){
    fmt.Println(values)
}

func addUsingVariadicParams(values ...int) (total int){
    for _,val := range values{
        total += val
    }
    return
}
