package main

import "fmt"

func main() {
	c := make(chan int)
	go squareNumbers(c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
		fmt.Println("Channel is closed")
	}

}

func squareNumbers(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i * i
	}
	close(c)
}
