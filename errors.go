package main

import (
	"errors"
	"fmt"
)

// В Go мы можем создавать пользовательские ошибки, реализовав error интерфейс в структуре.

type error interface {
	Error() string
}
type DivideByZero struct {
	message string
}

func (z DivideByZero) Error() string {
	return "Number can't divide by zero"
}

func divider(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, &DivideByZero{}
	} else {
		return num1 / num2, nil
	}
}

func main() {
	//for i := 0; i <= 5; i++ {
	//	result := 20 / i
	//	fmt.Println(result)
	//}
	// Мы можем обрабатывать ошибки с помощью:
	//
	// New()Функция
	//newError()

	// Errof() Функция
	//ErrofError()

	number1 := 9
	number2 := 0

	result, err := divider(number1, number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

// 1. Ошибка Go при использовании функции New()
func newError() {
	message := "Hello"
	myError := errors.New("Wrong Message")

	if message != "World" {
		fmt.Println(myError)
	}
	var name string
	fmt.Scanln(&name)
	err := checkName(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid name")
	}

}

func checkName(name string) error {
	newErr := errors.New("Invalid name ")

	if name != "Jemal" {
		return newErr
	}
	return nil
}

// 2.  Ошибка при использовании Errorf()

func ErrofError() {
	age := 14

	err1 := fmt.Errorf("%d is negative\nAge cannot be negative\n", age)

	if age < 0 {
		fmt.Println(err1)
	} else {
		fmt.Printf("Age: %d\n", age)
	}

	err2 := divide(2, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Println("Valid division")
	}

}

func divide(num1, num2 int) error {
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}
	return nil
}
