package main

import "fmt"

func main() {
	//var number int
	//fmt.Println("Enter your age : ")
	//fmt.Scan(&number)
	//
	//switch {
	//case number < 0:
	//	fmt.Println("Enter correct age")
	//	fmt.Scan(&number)
	//case number < 10 && number > 0:
	//	fmt.Println("You are child")
	//case number > 10 && number < 20:
	//	fmt.Println("You are too young")
	//case number < 50 && number > 20:
	//	fmt.Println("...")
	//default:
	//	fmt.Println("!!!!")
	//}

	var action string
	fmt.Println("Select action you need : (+, -, *, /)")
	fmt.Scan(&action)

	var a, b float64
	fmt.Println("Enter first number: ")
	fmt.Scan(&a)

	fmt.Println("Enter second number: ")
	fmt.Scan(&b)

	switch action {
	case "+":
		fmt.Println("Sum of numbers = ", a+b)
	case "-":
		fmt.Println("Sum of numbers = ", a-b)
	case "*":
		fmt.Println("Sum of numbers = ", a*b)
	case "/":
		fmt.Println("Sum of numbers = ", a/b)
	default:
		fmt.Println("Select correct action")

	}
}
