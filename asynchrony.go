package main

import (
	"fmt"
)

func main() {

	channel := make(chan int)

	go say("Hello", channel)
	// time.Sleep(2 * time.Second) channel ulanylsa time.Sleep hokman dal
	fmt.Println(<-channel)

	channel2 := make(chan int)
	go insert(channel2)
	for i := range channel2 {
		fmt.Println(i)
	}
	fmt.Println(<-channel2)
}

func say(greet string, channel chan int) {
	fmt.Println(greet)
	channel <- 7
}

func insert(channel chan int) {
	for i := 0; i <= 5; i++ {
		channel <- i
	}
	close(channel)
	// after close(channel) we cannot insert values to channel
}
