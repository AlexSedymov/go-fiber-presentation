package main

import "fmt"

func main() {
	// This line creates a new channel of type int
	c := make(chan int)
	// This line creates a new goroutine and calls the function squareNumbers
	go squareNumbers(c)
	// This line reads the channel and prints the value
	for i := 0; i < 5; i++ {
		// This line reads the channel and prints the value
		fmt.Println(<-c)
		fmt.Println("Channel is closed")
	}

}

// This function takes a channel of type int
func squareNumbers(c chan int) {
	for i := 0; i < 5; i++ {
		// This line writes the value of i * i to the channel
		c <- i * i
	}
	close(c)
}
