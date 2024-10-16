package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f func")
}

func f() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered in F func", r)
		}
	}()

	fmt.Println("Calling g func")
	g(0)
	fmt.Println("Returned normally from g func")
}

func g(num int) {
	if num > 3 {
		fmt.Println("Panicking!!!")
		panic(fmt.Sprintf("Num %v", num))
	}
	defer fmt.Println("Defer in g", num)
	fmt.Println("Printing in g", num)
	g(num + 1)
}
