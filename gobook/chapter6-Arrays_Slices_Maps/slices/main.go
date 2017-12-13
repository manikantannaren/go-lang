package main

import "fmt"

func main()  {
    // Slices don't have size in declaration unlike arrays.
    // This allows slices to expand
    // Slice works with underlying data arrays
    // The dimension of slice cannot be greater than the dimension of
    // underlying array

    // declares and creates a slice with an underlying array of length 5
    x := make([]float64, 5)

    //declares a slice of dimension 5 against an array of length 10
    x = make([]float64,5,10)

    //declare slice with index ranges - low and high
    arr:=[]float64{1,2,3,4,5,6,7,8,9,10,}
    //low index = 0, high index = 5 (i.e elements in positions <5)
    x = arr[0:5]
    fmt.Println(x)

    x = arr[2:7]
    //should print 3 to 7 (elements at position 2 to 6)
    fmt.Println(x)

    x = arr[:6]
    //shoudl print elements from positio 0 to 5
    fmt.Println(x)

    x = arr[6:]
    //shoudl print elements from position 6 to end
    fmt.Println(x)

    //append function
    slice1 := []int{1,2,3}
    //appends contents of slice1 and adds, elements 4,5 to the slice2
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)


    slice2 = append(slice2,6,7)
    fmt.Println(slice2)

    //copy function
    // slice2 has room for only two ele- ments only the first two elements of
    // slice1 are copied.
    slice2 = make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}
