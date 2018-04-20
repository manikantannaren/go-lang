//Channels are used to synchronize

package main

import (
	"fmt"
	"time"
)

type sint int

func producer(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping " + sint(i).String()

	}
}
func producer2(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping2 " + sint(i).String()

	}
}

func (i sint) String() string {
	return fmt.Sprintf("%d", i)
}

func consumer(c <-chan string, num string) {
	for {
		msg := <-c
		fmt.Println(num + "::" + msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)
	go producer(c)
	go producer2(c)
	go consumer(c, "1")
	go consumer(c, "2")
	var input string
	fmt.Scanln(&input)
}
