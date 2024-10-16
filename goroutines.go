package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// defer block
	//defer fmt.Println(1)   // 5
	//defer fmt.Println(2)   // 4
	//fmt.Println(add(2, 4)) // 1
	//deferValues()          // 2

	// goroutines block
	//runtime.GOMAXPROCS(1) // odnowremenno dine n<=NumCPU() sany goroutine isledyar
	//fmt.Println(runtime.NumCPU())

	//go showNumbers(100)

	//runtime.Gosched() // basga goroutine gecmek
	//time.Sleep(time.Second)

	//fmt.Println("Exit") // 3

	// Panic block
	//makePanic()

	t := time.Now()
	rand.Seed(t.UnixNano())

	go analyzeURL("youtube.com")
	analyzeURL("example.com")

	fmt.Printf("анализ завершен. истекшее время => %.2f seconds", time.Since(t).Seconds())

}
func showNumbers(number int) {
	for i := 0; i <= number; i++ {
		fmt.Println(i)
	}
}

func add(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()
	sum = x + y
	return
}

func deferValues() {
	//for i := 1; i <= 5; i++ {
	//	defer fmt.Println("First: ", i)
	//}
	//
	//for i := 1; i <= 5; i++ {
	//	defer func() {
	//		fmt.Println("Second: ", i)
	//	}()
	//}

	for i := 1; i <= 5; i++ {
		k := i
		defer func() {
			fmt.Println("Third: ", k)
		}()
	}

	for i := 1; i <= 5; i++ {
		defer func(k int) {
			fmt.Println("Fourth: ", k)
		}(i)
	}
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("Some panic")
	fmt.Println("Unreachable code")
}

func analyzeURL(url string) {
	for i := 1; i <= 5; i++ {
		zaderzhka := rand.Intn(500) + 500

		time.Sleep(time.Duration(zaderzhka) * time.Millisecond)

		fmt.Printf("Analyze <%s> - Step %d - zaderzhka %d ms\n", url, i, zaderzhka)

	}

}
