package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}

type Triangle struct {
	base, height float64
}

func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	rect := Rectangle{4, 8}
	triangle := Triangle{5, 5}

	printShapeArea(rect)
	printShapeArea(triangle)

}

func printShapeArea(s Shape) {
	fmt.Println("The Area of shape: ", s.area())
}
