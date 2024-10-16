package main

import (
	"fmt"
	"unsafe"
)

const nameJ = "Jemal"

func main() {

	var a string = "Hello world"
	var b bool
	fmt.Println(a)
	fmt.Println(b)
	b = true
	fmt.Printf("Type: %T Value: %v\n ", a, a)
	fmt.Printf("Type: %T Value: %v\n ", b, b)
	// mesto var dostatochno :=
	c := 26
	fmt.Println(c)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Println(nameJ)
	fmt.Printf("Type: %T Value: %v\n", nameJ, nameJ)

	m := "Nick"
	n := fmt.Sprintf("Hello %s", m)
	//_ = n
	fmt.Println(n)

	//Operations

	var num1 int64 = 30
	num2 := 25
	fmt.Println(num1 + int64(num2))

	var x uint8 = 1
	var y uint64 = 1
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y))

	at := "Nick"
	fmt.Println(at)

	fmt.Printf("Hi %s\n", at)

	// simple func example1

	sqr := func(x int) int {
		return x * x
	}

	sqrA(sqr)

}
