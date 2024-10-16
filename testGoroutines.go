package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	message := make(chan string)

	go numbersChannel(numbers)
	go stringChannel(message)

	select {
	case firstChannel := <-numbers:
		fmt.Println("Channel data: ", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel data: ", secondChannel)

	default:
		fmt.Println("Wait")
	}

}
func numbersChannel(ch chan int) {
	//time.Sleep(2 * time.Second)
	ch <- 15
}

func stringChannel(ch chan string) {
	//time.Sleep(2 * time.Second)
	ch <- "Message"
}
