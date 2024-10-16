package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fruitsString := "12 3 2  1"

	fruit := strings.Split(fruitsString, " ")
	fmt.Printf("Type %T, value %v\n", fruit, fruit)
	//output : {"1", "2", "i", "2"}

	//fruit := strings.SplitN(fruitsString, " ", 3) // jemi n bolege bolup beryar
	//fmt.Printf("Type %T, value %#v\n", fruit, fruit)
	// output : {"1", "2", "i 2"}

	// Аналогично Split(), SplitAfter()разбивает исходную строку,
	//но оставляет разделители в конце каждой подстроки.

	//fruit := strings.SplitAfter(fruitsString, " ")
	//fmt.Printf("Type %T, value %#v\n", fruit, fruit)
	// output : {"1 ", "2 ", "i ", "2"}

	// разделяет только первые N подстроки.
	//fruit := strings.SplitAfterN(fruitsString, " ", 3)
	//fmt.Printf("Type %T, value %#v\n", fruit, fruit)
	// output : {"1 ", "2 ", "i 2"}

	// разделяет по пробелам и новым строкам
	//fruit := strings.Fields(fruitsString)
	//fmt.Printf("Type %T, value %#v\n", fruit, fruit)
	//// output : {"1", "2", "i", "2"}

	// разделяет с помощью регулярного выражения
	//fruit := regexp.MustCompile("[a-z]").Split(fruitsString, -1)
	//fmt.Printf("Type %T, value %#v\n", fruit, fruit)
	//
	//

	fmt.Println("Enter your number: ")
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = input[:len(input)-1]
	fmt.Printf("%T => %#v\n", input, input)
	splitString(input)
}

func splitString(str string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var result float64

	fruits := strings.Split(str, " ")
	for _, val := range fruits {
		if val == "" {
			continue
		}
		//fmt.Printf("Type %T value %#v\n", val, val)
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Printf("%T => %#v\n", num, num)
			panic("Enter only number")
		}
		fmt.Printf("Type %T value %#v\n", num, num)
		result += num
	}
	fmt.Println("Sum of numbers: ", result)
}
