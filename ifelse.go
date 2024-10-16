package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("a : ")
	fmt.Scan(&a)

	fmt.Println("b : ")
	fmt.Scan(&b)

	fmt.Println("c : ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("x1=", x1, " x2= ", x2, " and D= ", D)
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Println("x= ", x, " and D= ", D)
	} else if D < 0 {
		fmt.Println("D= ", D)
	}

	a2 := 0

	// simple if, full syntax

	if a2 < 0 {
		fmt.Println(a, "San noldan kici (full)")
	}

	// simple if, short syntax

	if minus := number(a2); minus == true {
		fmt.Println(a, " San noldan kici (short)")
	}

	// if...else

	if a2 < 0 {
		fmt.Println("San noldan kici")

	} else if a2 > 0 {
		fmt.Println("San noldan uly")
	} else {
		fmt.Println("San nola den")
	}

	// or   || / and &&

	age := 25

	if age > 6 && age < 18 {
		fmt.Println("You are in school")
	} else {
		fmt.Println("You are not in school")
	}

	if age == 16 || age == 25 || age == 40 {
		fmt.Println("You have to get new passport")
	}

	if !Age(age) {
		fmt.Println("Your age is ", age, ", you are an adult!")
	}

}
func number(a2 int) bool {
	return a2 < 0
}

func Age(age int) bool {
	return age < 18
}
