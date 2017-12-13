package main
import "fmt"

func  main()  {
    factorial := recursiveFuncForFactorial(4)
    fmt.Println("Factorial of 4: ",factorial)
}

func recursiveFuncForFactorial(num int) (factorial int){
    if num == 0{
        factorial = 1
        return
    }

    return num* recursiveFuncForFactorial(num-1)
}
