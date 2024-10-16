package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func changeArr(arr [3]int) [3]int {
	arr[1] = 7
	return arr
}

func main() {

	//// simple array example 1
	//var array1 [3]string
	//array1[0] = "first"
	//array1[1] = "second"
	//array1[2] = "third"
	////fmt.Println(array1)
	//
	//// short syntax
	//
	//array2 := [3]string{"first", "second", "third"}
	//fmt.Println(array2, len(array2))
	//fmt.Printf("%T, %v\n", array2, array2)
	//
	//for i := 0; i < len(array1); i++ {
	//	//fmt.Println(i, ") ", array1[i])
	//}
	//
	//numbers := [5]float64{12, 13, 11, 16, 31}
	//var sum float64 = 0
	//for j := 0; j < len(numbers); j++ {
	//	sum = sum + numbers[j] // 	sum += numbers[j]
	//}
	//var result float64 = sum / float64(len(numbers))
	//fmt.Println("Orta arifmetik bahasy : ", math.Round(result))
	//
	//// two-dimensional array;  multidimensional array (mnogomernyy massiw)
	//
	//array3 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	//fmt.Println(array3)
	//fmt.Println(array3[2])
	//fmt.Println(array3[2][1])
	//
	//// slice ( srezy) - dynamic arrays or olcegsiz massiw
	//slyce := []int{2, 3, 4, 1, 9, 0}
	//slyce = append(slyce, 5) // add element to end
	//fmt.Println(slyce)
	//
	//slyce[4] = 8 // change element
	//fmt.Println(slyce)
	//
	//sort.Ints(slyce) // sort elments min to max
	//fmt.Println(slyce)
	//
	////
	//slyce2 := []string{"Mark", "John", "Nick", "Anna"}
	//slyce2 = append(slyce2, "Z")
	//sort.Strings(slyce2) // sort string elements
	//fmt.Println(slyce2)
	//
	//// advanced array iteration
	//
	//array4 := []int{2, 1, 3, 2, 4, 56, 7, 8}
	//
	//for _, element := range array4 { // _, here i,
	//	fmt.Println(element)
	//}
	//
	//// an integer array
	//nmbrs := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	//
	//// create slice from an array
	//sliceNumbers := nmbrs[4:7]
	//
	//fmt.Println(sliceNumbers)
	//
	//// create two slices
	//evenNumbers := []int{2, 4, 5, 3, 1}
	//oddNumbers := []int{1, 3}
	//
	//// add elements of oddNumbers to evenNumbers
	//evenNumbers = append(evenNumbers, oddNumbers...)
	//fmt.Println("Numbers:", evenNumbers)
	//
	//// copy elements of primeNumbers to numbers
	//copy(evenNumbers, oddNumbers)
	//
	//// print numbers
	//fmt.Println("Numbers:", oddNumbers)

	// Exampls from This is it

	var intArr [3]int
	fmt.Printf("Type: %T, Value: %#v\n", intArr, intArr)
	intArr[0] = 2
	intArr[2] = 6
	fmt.Printf("Type: %T, Value: %#v\n", intArr, intArr)
	people := [2]Person{
		{
			"Person1",
			78,
		},
		{
			"Person2",
			45,
		},
	}
	fmt.Printf("Type: %T, value: %#v\n", people, people)

	stringArr := [...]string{"First", "Second", "Third", "Fourth"}
	fmt.Printf("Type: %T, value: %#v\n", stringArr, stringArr)

	// functions len {length} and cap {capacity}
	fmt.Printf("Length: %d, capacity: %d\n", len(stringArr), cap(stringArr))

	for i := 0; i < len(stringArr); i++ {
		fmt.Printf("Index = %d, value = %s\n", i, stringArr[i])
	}
	fmt.Println()

	for index, value := range stringArr {
		fmt.Printf("Index : [%d], Value = \"%s\" \n", index, value)
	}
	fmt.Println()

	for _, num := range intArr {
		fmt.Printf("The number is : %d\n", num)
	}
	fmt.Println()
	newIntArr := changeArr(intArr)
	fmt.Printf("intArr type: %T, intArr values: %d\n", intArr, intArr)
	fmt.Printf("new intArr type: %T, new intArr values: %d\n", newIntArr, newIntArr)
}
