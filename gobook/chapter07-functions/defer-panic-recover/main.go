package main

import "fmt"

func main()  {
    // defers moves invocation to end of function
    // Useful for closing streams for example at the end
    // defer works even if intermediary steps throw a panic
    // defer is like finally
    defer fmt.Println(recursiveFuncForFactorial(5))
    defer func (){
        str:=recover() //catch block
        fmt.Println(str)
        fmt.Println(recursiveFuncForFactorial(6))
    }()
    panic("some panic")

}

func recursiveFuncForFactorial(num int) (factorial int){
    if num == 0{
        factorial = 1
        return
    }

    return num* recursiveFuncForFactorial(num-1)
}
