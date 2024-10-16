package main

import "fmt"

func main() {

	for a := 1; a <= 10; a++ {
		//fmt.Printf("Hello %d \n", a)
	}

	for b := 1; b < 20; {
		b += b
		fmt.Println(b)
	}
Label1:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			//fmt.Println(i, " x ", j, " = ", i*j)
			if i >= 5 {
				break Label1
			}
		}
		//	fmt.Println("\n")

	}

	// usage while
	m := 0
	for m < 10 {
		//fmt.Println(m)
		m++
	}

	massiw := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(massiw); i++ {
		//fmt.Println(massiw[i])
	}

	//for index, element := range massiw {
	//	fmt.Printf("Index: %d Element: %d\n", index, element)
	//}
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, a2 := range matrix {
		for _, b2 := range a2 {
			fmt.Printf("%d ", b2)
		}
		fmt.Println()
	}
}
