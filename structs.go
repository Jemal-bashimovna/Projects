package main

import "fmt"

// Example declare (obyawit) struct
type Students struct {
	name string
	age  int
	mark float64
}

// Example declare (obyawit) struct

type rectangle struct {
	length  int
	breadth int
}
type rectangle1 func(int, int) int

func change(a *rectangle2) {
	a.color = "green"
	a.length = 10
}

type rectangle2 struct {
	length  int
	breadth int
	color   string
	rect    rectangle1
}

func main() {
	var person1 Students
	person1 = Students{"John", 26, 57.7}

	fmt.Println(person1)
	fmt.Println("Name: ", person1.name, "age: ", person1.age, "mark: ", person1.mark)

	person2 := Students{"Mark", 24, 68.4}
	fmt.Println(person2)

	rect1 := rectangle{23, 15}

	fmt.Println("Length: ", rect1.length, "\nBreadth: ", rect1.breadth)
	sqrRect1 := square(rect1.length, rect1.breadth)
	fmt.Println("Square:", sqrRect1)
	prmtrRect1 := perimeter(rect1.length, rect1.breadth)
	fmt.Println("Perimeter:", prmtrRect1)

	result := rectangle2{
		6,
		8,
		"red",
		func(length int, breadth int) int {
			return length * breadth
		},
	}
	change(&result)
	fmt.Println("Color of rectangle: ", result.color)
	fmt.Println("Area of rectangle: ", result.rect(result.length, result.breadth))

}
func square(a, b int) int {
	return a * b
}

func perimeter(a, b int) int {
	return (a + b) * 2
}
