package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type numberInterface interface {
	Sum() int
	Multiply() int
	Division() float64
	Subtraction() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n Numbers) Subtraction() int {
	return n.num1 - n.num2
}

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	// implementation of interface Runner
	return fmt.Sprintf("Human %s is runnig", h.Name)

}

func main() {
	var i numberInterface
	numbers := Numbers{45, 20}
	i = numbers
	fmt.Printf("Sum of numbers: %d + %d = %d\n", numbers.num1, numbers.num2, i.Sum())
	fmt.Printf("Multiply of numbers: %d x %d = %d\n", numbers.num1, numbers.num2, i.Multiply())
	fmt.Printf("Division of numbers: %d : %d = %.2f\n", numbers.num1, numbers.num2, i.Division())
	fmt.Printf("Subtraction of numbers: %d - %d = %d\n\n", numbers.num1, numbers.num2, i.Subtraction())

	//This is it
	var runner Runner
	fmt.Printf("Type: %T value: %#v\n", runner, runner)

	// interfaceValue = nil

	if runner == nil {
		fmt.Println("Runner is nil")
	}

	// implementation
	//runner = int64(68)

	var unnamedRunner *Human
	fmt.Printf("Type: %T value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Type: %T value: %#v\n", runner, runner)

	//               Concrete Type
	//              /
	// interfaceValue
	//              \
	//               nil

	if runner == nil {
		fmt.Println("Runner is nil")
	}

	namedRunner := &Human{"John"}
	fmt.Printf("Type: %T value: %#v\n", namedRunner, namedRunner)

	//               Concrete Type
	//              /
	// interfaceValue
	//              \
	//               Concrete Value

	runner = namedRunner
	fmt.Printf("Type: %T value: %#v\n", runner, runner)

	fmt.Println(runner.Run())

	// empty interface
	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("Type: %T value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = namedRunner
	fmt.Printf("Type: %T value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(12)
	fmt.Printf("Type: %T value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("Type: %T value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = "Moscow"
	fmt.Printf("Type: %T value: %#v\n", emptyInterface, emptyInterface)

	var nilInterface interface{}
	fmt.Println(nilInterface)
}
