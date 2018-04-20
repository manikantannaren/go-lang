package main
import "fmt"

func main()  {
    //x is a map of string keys and int values
    //Maps are unordered collections of keys and values
    /*
    var x map[string]int

    x["ten"] = 10
    fmt.Println(x) //will throw run time exception since the map is declared but not initialized - a.k.a. NullPointerException
    */

    x := make(map[string]int)
    x["ten"] = 10
    fmt.Println(x)

    x["one"] = 1
    x["two"] = 2
    fmt.Println(x)

    delete(x, "one")
    fmt.Println(x)

    //no value for a key - contains(key) == false
    value, ok := x["Un"]


    if ok{
        fmt.Println(value, ok)
    }else{
        fmt.Println("no key ", "un")
    }

    value,ok = x["two"]
    if ok{
        fmt.Println(value, ok,"two")
    }else{
        fmt.Println("no key ", "un")
    }
}
