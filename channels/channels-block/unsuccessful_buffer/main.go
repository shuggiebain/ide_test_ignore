package main

import "fmt"

func main() {
	c := make(chan int, 1) //buffer channel allows one value to sit in there

	c <- 42
	c <- 43

	fmt.Println(<-c)

}
