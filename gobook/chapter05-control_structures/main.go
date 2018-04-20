package main
import "fmt"
func main() {
    //can print a series of values
    fmt.Println(`1
2
3
4
5
6
7
8
9
10`)

//using for loop
//Java while and for loop can be done with for in golang
//For loop like Java while
    fmt.Println("Like Java while")
    i:= 1
    for i<=10{
        fmt.Println(i)
        i++
    }

    fmt.Println("Proper For loop")
    for i:=1; i <=10; i++{
        fmt.Println(i)
    }

    fmt.Println("IF conditional print")
    for i:=1; i <=10; i++{
        if(i % 2 == 0){
            fmt.Println(i, "Even")
        }else{
            fmt.Println(i, "Odd")
        }
    }

    fmt.Println("Switch statement")
    for i:=1; i <=10; i++{
        switch i{
            case 1: fmt.Println("One")
            case 2: fmt.Println("Two")
            default: fmt.Println("Some number : ",i)
        }
    }

}
