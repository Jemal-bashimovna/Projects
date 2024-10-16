package main

import "fmt"

func sum(a int, b int) (int, int, int, float32) {
	first := a + b
	second := a - b
	third := a * b
	fourth := float32(a / b)
	return first, second, third, fourth
}

func test(someFunc func(int) int) {
	fmt.Println(someFunc(6))
}

func test1(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {

	x, y, z, w := sum(3, 9)
	fmt.Println(x, y, z, w)

	// func inside of func

	square := func(x int) int {
		return x * x
	}

	test(square)

	test1("Hello ")()

	//first, second := 2, 6
	//
	//aFunc := func(x, y int) int {
	//	return x + y
	//}
	//
	//bFunc := func(x, y int) int {
	//	return x - y
	//}
	//
	//fmt.Println(calculate(first, second, aFunc))
	//fmt.Println(calculate(second, first, bFunc))
	//

	divBy2 := createDivider(2)
	fmt.Println(divBy2(80))

	// course of dollar
	dol := 40
	Value := func() int {
		return dol
	}
	fmt.Println(Value())
	dol = 90

	fmt.Println(Value())

	// recursive function ex1
	count(5)

	// recursive function ex2
	//var f int
	//fmt.Scan(&f)
	//fmt.Println(suM(f))

	// recursion for factorial of number
	var factorial int
	fmt.Scan(&factorial)
	fmt.Println(fact(factorial))

}

func createDivider(divider int) func(x int) int {
	divFunc := func(x int) int {
		return x / divider
	}
	return divFunc
}

func sqrA(sqrB func(int) int) {
	fmt.Println(sqrB(7))
}

// recursive function ex1
func count(x int) {
	if x >= 0 {
		fmt.Println(x)
		count(x - 1)
	} else {
		fmt.Println("Time is over")
	}
}

// recursive function ex2

func suM(y int) int {
	if y == 0 {
		return 0
	} else {
		return y + suM(y-1)
	}
}

// recursion for factorial of number

func fact(i int) int {
	if i == 0 {
		return 1
	} else {
		return i * fact(i-1)
	}
}
