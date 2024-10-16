package main

import (
	"fmt"
	"io"
	"os"
)

/*В Go мы используем операторы defer, panic и recovery для обработки ошибок.
Мы используем defer для задержки выполнения функций, которые могут вызвать ошибку.
Оператор panic немедленно завершает программу и recover используется
 для восстановления сообщения во время паники */

// Мы используем defer оператор, чтобы предотвратить выполнение функции,
//пока не будут выполнены все остальные функции.

func main() {
	//defer fmt.Println(3)
	//fmt.Println(1)
	//fmt.Println(2)

	//fmt.Println("При использовании нескольких операторов defer, последний оператор defer будет выполнен первым.")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)

	//	Мы используем оператор паники, чтобы немедленно завершить
	//	выполнение программы. Если наша программа достигает точки,
	//	где ее невозможно восстановить из-за
	//	некоторых серьезных ошибок, лучше всего использовать панику.

	fmt.Println("Something is wrong")
	panic("Ending the program")
	fmt.Println("Unreachable code")
	division(30, 3)
	division(4, 0)
	division(55, 5)

	//	Оператор паники немедленно завершает программу. Однако иногда
	//	программе может быть важно завершить свое выполнение и получить требуемые результаты.
	//  В таких случаях мы используем оператор recover

	//fmt.Println(copyFile("text2.txt", "text1.txt"))

	fmt.Println(c())

	// Recover полезен только внутри отложенных функций.
	//Во время обычного выполнения вызов recovery вернет
	//nil и не будет иметь никакого другого эффекта.
}
func c() (i int) {
	defer func() {
		i++
	}()
	return 3
}

func division(num1, num2 int) {
	defer handlePanic()

	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

// recover function to handle panic
func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}
}

func copyFile(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
