package main

import "fmt"

/*
	1. create method :

type T struct{}

func (receiver T) methodName(){}
func (receiver *T) methodName(){}
*/

type OurType string

type Kwadrat struct {
	tarap int
}

func (t OurType) Greet() {
	fmt.Println("Hello world")
}

func (a Kwadrat) Perimetr() {
	fmt.Printf("Tipi : %T ; bahasy: %#v\n", a, a)
	fmt.Printf("Kwadratyn perimetri: %d\n", a.tarap*4)
}

func (a *Kwadrat) Kopeldiji(k int) {
	fmt.Printf("Tipi : %T ; bahasy : %#v; kopeldiji : %d\n", a, a, k)
	a.tarap = a.tarap * k
	fmt.Printf("%d x  %#v\n", k, a)
}

func main() {
	kwadrat := Kwadrat{tarap: 3}
	//pointerKwadrat := &kwadrat
	//pointerKwadrat.Perimetr()
	//
	//pointerKwadrat.Kopeldiji(5)

	kwadrat.Kopeldiji(4)
	kwadrat.Kopeldiji(2)

}
