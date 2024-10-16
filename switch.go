package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minI = 1
	maxI = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(maxI-minI) + minI
	// if example
	if value == 1 {
		fmt.Println("Number is 1")
	} else if value == 2 || value == 3 {
		fmt.Println("Number is 2 or 3")
	} else if value == four() {
		fmt.Println("Number is 4")
	} else {
		fmt.Println("number is 5")
	}

	// simple switch - 1 example
	switch value {
	case 1:
		fmt.Println("Number is 1")
	case 2, 3:
		fmt.Println("Number is 2 or 3")
	case four():
		fmt.Println("Number is 4")
	default:
		fmt.Println("Number is 5")
	}

	// switch - 2

	switch num := rand.Intn((maxI - minI) + minI); num {
	case 1:
		fmt.Println("Number is 1")
	case 2, 3:
		fmt.Println("Number is 2 or 3")
	case four():
		fmt.Println("Number is 4")
		fallthrough
	case 10:
		fmt.Println("Strange thing")
	default:
		fmt.Println("Number is 5")

	}

	// switch - 3

	switch {
	case value > 2:

		fmt.Printf("Value %d greater than 2\n", value)
	case value < 2:
		fmt.Printf("Value %d smaller than 2\n", value)
	default:
		fmt.Printf("Value equals 2")
	}

}

func four() int {
	fmt.Println("The function getFour is called")
	return 4
}
