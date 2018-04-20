package main
import "fmt"

/*similar to java anonymous classes*/
func main()  {
    addUsingClosure := func (values ...int) (total int, n int){
        n = len(values)
        for _,val:=range values{
            total += val
        }
        return
    }
    slicedValues := []int{5,15,25,35}
    sum,n := addUsingClosure(slicedValues...)
    fmt.Println("Number of elements",n, " and Total ",sum)

    generator := seqGeneratorFunction()
    for index := 0;  index < 10; index++ {
        fmt.Println("Using generator : ", generator())
    }

    evenGenerator := oddOrEvenNumberGenerator(false)
    for index := 0;  index < 10; index++ {
        fmt.Println("Even numbers using generator : ", evenGenerator())
    }

    oddGenerator := oddOrEvenNumberGenerator(true)
    for index := 0;  index < 10; index++ {
        fmt.Println("Odd numbers using generator : ", oddGenerator())
    }

}

func seqGeneratorFunction() func () uint{
    i:= uint(0)
    return func () (nextNum uint){
        i++
        nextNum = i
        return
    }
}

func oddOrEvenNumberGenerator(oddNumbers bool) func () int{
    odd := 1
    even := 0
    return func () (nextNum int){
        if oddNumbers{
            nextNum = odd
            odd += 2
        }else{
            nextNum = even
            even += 2
        }
        return
    }
}
