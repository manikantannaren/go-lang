package main

import "fmt"

func main() {
	x := 5
	zeroFunc(&x) //pass address
	fmt.Println(x)

	xPtrWithNew := new(int)
	withNewFunc(xPtrWithNew)
	fmt.Println(xPtrWithNew) //prints mem address
	fmt.Println(*xPtrWithNew)
}

func withNewFunc(ptr *int) {
	*ptr = 5
}

func zeroFunc(xptr *int) { //*int is pointer
	*xptr = 0 //change value at &x
}
