package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	// "fmt" package
	fmt.Println("hello")
	//fmt.Println("the random number is: ", rand.Intn(20))
	//fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))

	// create variables
	var name string
	var age int

	// take name and age input
	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age)
	//fmt.Scanln(&name, &age)
	//fmt.Scanf(&name, &age)
	// take name and age input using format specifier
	//fmt.Println("Enter your name and age:")
	//fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	var number int

	//// take input value
	//fmt.Scan(&number)
	//
	//// print using Println
	//fmt.Println("Number is", number)
	//
	//fmt.Print("Using Print")
	//fmt.Println("Using Println")

	fmt.Scanf("%d", &number)   // Input: 10
	fmt.Printf("%d\n", number) // Output: 10

	// "math" package

	// find the square root
	fmt.Println(math.Sqrt(25)) // 5

	// find the cube root
	fmt.Println(math.Cbrt(27)) // 3

	// find the maximum number
	fmt.Println(math.Max(21, 18)) // 21

	// find the minimum number
	fmt.Println(math.Min(21, 18)) // 18

	// find the remainder
	fmt.Println(math.Mod(5, 2)) // 1

	// "strings" package

	// convert the string to lowercase
	lower := strings.ToLower("GOLANG STRINGS")
	fmt.Println(lower)

	// convert the string to uppercase
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

	// create a string array
	stringArray := []string{"I love", "Go Programming"}

	// join elements of array with space in between
	joinedString := strings.Join(stringArray, " ")
	fmt.Println(joinedString)

	//Take Input of Float and Boolean Type
	var temperature float32
	var sunny bool

	// take float input
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%g", &temperature)

	// take boolean input
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)

	f

}
