package main

import "fmt"

// Каналы в Go служат средством общения горутин друг с другом.
func main() {
	//// create channel of integer type
	//channelName := make(chan int)
	//fmt.Printf("Type %T\nValue %v\n", channelName, channelName)
	//
	////	 значение канала — это адрес памяти, который
	////	 действует как среда, через которую горутины
	////	 отправляют и получают данные для взаимодействия.
	//
	//numbers := make(chan int)
	//message := make(chan string)
	//go channelData(numbers, message)
	//
	//fmt.Println("Channel Data: ", <-numbers)
	//fmt.Println("Channel Data: ", <-message)
	//
	//ch := make(chan string)
	//
	//go sendData(ch)
	//
	//fmt.Println(<-ch)

	chann := make(chan string)
	go receiveData(chann)

	fmt.Println("No data. Receive operation blocked")

	chann <- "Data Received. received operation successful"
	go receiveData(chann)
}

func channelData(number chan int, message chan string) {

	number <- 24
	message <- "Hello world"
}

func sendData(ch chan string) {
	ch <- "Send operation Successful"
	fmt.Println("Send operation blocked")
}

func receiveData(ch chan string) {
	fmt.Println(<-ch)
}
