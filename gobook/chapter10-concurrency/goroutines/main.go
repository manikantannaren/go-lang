package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//calling a function using the keyword 'go' runs it parallel to main/calling function
	go printValues(0, 2)
	fmt.Println("main continuing")
	var input string
	fmt.Scanln(&input) //pause main so that numbers get printed else, it will quit

	//go run multiple times - each go routine seems to run one after the other
	for i := 0; i < 3; i++ {
		go printValues(i, i+1)
	}
	fmt.Println("main continuing: ", 2)
	fmt.Scanln(&input) //pause main so that numbers get printed else, it will quit

	for i := 0; i < 3; i++ {
		// the delay in called function will make it seem like the routines are running in parallel
		go printValuesWithDelay(i, i+1)
	}
	fmt.Println("main continuing: ", 3)
	fmt.Scanln(&input) //pause main so that numbers get printed else, it will quit

}

func printValues(callerID int, numberOfValues int) {
	for i := 0; i < numberOfValues; i++ {
		fmt.Println("printing values for caller: ", callerID, " Value : ", i)
	}
}

func printValuesWithDelay(callerID int, numberOfValues int) {
	for i := 0; i < numberOfValues; i++ {
		fmt.Println("printing delayed values for caller: ", callerID, " Value : ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
