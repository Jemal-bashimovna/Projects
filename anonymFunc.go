package main

import "fmt"

func main() {

	// anonymous function is unnamed func
	var a = func() {
		fmt.Println("string")
	}

	a() // function call

	// anonymous function with arguments
	//var num1, num2 int
	//
	//fmt.Println("Enter number 1: ")
	//fmt.Scan(&num1)
	//fmt.Println("Enter number 2: ")
	//fmt.Scan(&num2)
	//
	//a1 := func(num1, num2 int) (int, int) {
	//	result1 := num1 + num2
	//	result2 := num1 - num2
	//	return result1, result2
	//}
	//
	//fmt.Println(a1(num1, num2)) // function call

	// Анонимная функция как аргументы других функций:

	// anonymous function to calculate sum of numbers
	s := func(n1, n2 int) int {
		return n1 + n2
	}

	// call function

	result := sqr(s(2, 4))
	fmt.Printf("Result of (num1+num2)^2 = %d\n", result)

	//Возврат анонимной функции :
	// call function that returns an anonymous function
	a2 := func1()
	fmt.Println(a2())

	// Вложенная функция
	// call outer function
	greet("Jemal")

	// Closers function icindaki funksiyalara function dashyndan hem yuzlenip bolar yaly edyar
	greet3 := greet2()
	fmt.Println(greet3())

	// печать нечетных чисел (odd numbers) с использованием closer
	oddNum := odd() // call function

	fmt.Println(oddNum())
	fmt.Println(oddNum())
	fmt.Println(oddNum())

	oddNum2 := odd() // call function again
	fmt.Println(oddNum2())
	fmt.Println(oddNum2())
	fmt.Println(oddNum2())
}

// Анонимная функция как аргументы других функций:
// regular function to calculate square of numbers

func sqr(n int) int {
	return n * n

}

// Возврат анонимной функции :
// function that returns an anonymous function
func func1() func() int {
	san := 20
	return func() int {
		return san + 20
	}
}

// Вложенная функция
// create outer function
func greet(name string) {
	// create inner function
	var personGreet = func() {
		fmt.Println("Hi", name)
	}
	personGreet() // call inner function
}

// Closers

func greet2() func() string {
	// variable defined outside the inner function
	person := "Mike"
	return func() string {
		person = "Hello " + person
		return person
	}

}

// печать нечетных чисел (odd numbers) с использованием closer
func odd() func() int {
	k1 := 1
	return func() int {
		k1 = k1 + 2
		return k1
	}
}
