package main

import "fmt"

func main() {

	//	Simple Slices
	//var defaultSlice []int
	//fmt.Printf("Type %T, value %#v\n", defaultSlice, defaultSlice)
	//fmt.Printf("Length = %d, Capacity = %d\n", len(defaultSlice), cap(defaultSlice))
	//fmt.Println()
	//
	//stringSlice := []string{"First", "Second"}
	//fmt.Printf("Type %T, value %#v\n", stringSlice, stringSlice)
	//fmt.Printf("Length = %d, Capacity = %d\n", len(stringSlice), cap(stringSlice))
	//fmt.Println()
	//
	//sliceByMake := make([]int, 3, 4)
	//fmt.Printf("Type %T, value %#v\n", sliceByMake, sliceByMake)
	//fmt.Printf("Length = %d, Capacity = %d\n", len(sliceByMake), cap(sliceByMake))
	//fmt.Println()
	//
	//sliceByMake = append(sliceByMake, 1, 2, 4)
	//fmt.Printf("Type %T, value %#v\n", sliceByMake, sliceByMake)
	//fmt.Printf("Length = %d, Capacity = %d\n", len(sliceByMake), cap(sliceByMake))
	//fmt.Println()
	//// capacity basdakydan 2 esse ulalyar
	//
	//sliceByMake = append(sliceByMake, 1, 5, 6, 78, 0)
	//fmt.Printf("Type %T, value %#v\n", sliceByMake, sliceByMake)
	//fmt.Printf("Length = %d, Capacity = %d\n", len(sliceByMake), cap(sliceByMake))
	//fmt.Println()
	//
	//for i, v := range sliceByMake {
	//	fmt.Printf("Index [%d] value = %d\n", i, v)
	//}

	//slices()
	//convertToArrayPointer()
	//passToFunction()
	sliceWithNew()
	//getSlice()
	//copySlice()
	//deleteElementT()
}

func slices() {
	showAllElements(1, 3, 5)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{9, 8, 7, 6, 5}
	secondSlice := []int{4, 3, 2, 1}

	showAllElements(firstSlice...)                 // (9,8,7,6,5)
	newSlice := append(firstSlice, secondSlice...) //(firstSlice, 4,3,2,1)
	fmt.Printf("Type %T value %#v\n", newSlice, newSlice)
}

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Println("Value: ", val)
	}
	fmt.Println()
}

func convertToArrayPointer() {
	iSlice := []int{1, 2}
	fmt.Printf("Type %T value %#v\n", iSlice, iSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(iSlice), cap(iSlice))
	fmt.Println()

	intArray := (*[2]int)(iSlice) // *[2] olcegi iSlice olcegi bilen gabat gelmeli
	fmt.Printf("Type %T value %#v\n", intArray, intArray)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intArray), cap(intArray))
	fmt.Println()
}

func passToFunction() {
	iSlice := []int{1, 2, 3}
	fmt.Printf("Type %T value %#v\n", iSlice, iSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(iSlice), cap(iSlice))
	fmt.Println()

	// Structure of Slice:
	/* type _slice struct {
		{
		elements unsafe.Pointer // ukazatel na massiw
	len int // kolichestwo elementow
	cap int  // tekushshaya wmestimost
	}
	}

	*/

	changeValue(iSlice) // Slice funksiva ugradanymyzda,massiwe bolan pointer ugradylyar (return hokman dal value uytgedilende	)
	fmt.Printf("Type %T value %#v\n", iSlice, iSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(iSlice), cap(iSlice))
	fmt.Println()

	newSlice := append(iSlice, 3)
	fmt.Printf("Type %T value %#v\n", newSlice, newSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(newSlice), cap(newSlice))
	fmt.Println()

	// taze peremena doretmezden append edip bolyar
	iSlice = append(iSlice, 0)
	fmt.Printf("Type %T value %#v\n", iSlice, iSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(iSlice), cap(iSlice))
	fmt.Println()

	// funksiya arkaly taze value gosmak
	newSlice2 := appendValue(iSlice)
	fmt.Printf("Type %T value %#v\n", newSlice2, newSlice2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(newSlice2), cap(newSlice2))
	fmt.Println()
}
func changeValue(array []int) {
	array[2] = 10
}

func appendValue(slice []int) []int {
	slice = append(slice, 9)
	//fmt.Printf("Type %T value %#v\n", slice, slice)
	//fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	//fmt.Println()
	return slice
	// 	slice (pointeri, len cap uytgedi )atly taze slice doredya
}

func sliceWithNew() {
	slicePointer := new([]int)
	fmt.Printf("Type %T value %#v\n", slicePointer, *slicePointer)
	fmt.Printf("Length : %d, Capacity : %d\n\n", len(*slicePointer), cap(*slicePointer))

	// slicePointer=append(slicePointer, 1,2) - basga ululyga dakmazdan value goshup bolano new() bilen doredilende

	newSlice := append(*slicePointer, 3, 4)
	fmt.Printf("Type %T value %#v\n", newSlice, newSlice)
	fmt.Printf("Length : %d, Capacity : %d\n\n", len(newSlice), cap(newSlice))

}

func getSlice() {

	// RESLICING
	intArr := [...]int{0, 1, 2, 3, 4}
	fmt.Printf("Type %T value %#v \n\n", intArr, intArr)

	intSlice := intArr[1:3] // slice from array
	// index=1 den index=3 cenli elementleri alyar, (index[3] alanok), capacity index[1] den sonky elementleirn sanyna gora doredilyar
	fmt.Printf("Type %T value %#v \n", intSlice, intSlice)
	fmt.Printf("Length : %d, Capacity : %d\n\n", len(intSlice), cap(intSlice))

	fullSlice := intArr[:] // [0:5]
	fmt.Printf("Type %T value %#v \n", fullSlice, fullSlice)
	fmt.Printf("Length : %d, Capacity : %d\n\n", len(fullSlice), cap(fullSlice))

	//slice from slice
	intSlice2 := fullSlice[:4]
	fmt.Printf("Type %T value %#v \n", intSlice2, intSlice2)
	fmt.Printf("Length : %d, Capacity : %d\n\n", len(intSlice2), cap(intSlice2))

	intArr[1] = 300 // sonky slice hemmesi baslangyc berlen array salgylanyar
	fmt.Printf("Type %T value %#v \n", intArr, intArr)
	fmt.Printf("Type %T value %#v \n", intSlice, intSlice)
	fmt.Printf("Type %T value %#v \n", fullSlice, fullSlice)
	fmt.Printf("Type %T value %#v \n", intSlice2, intSlice2)
}

func copySlice() {
	slice1 := make([]string, 0, 2)
	slice2 := []string{"First", "Second", "Third"}

	fmt.Println("Copied", copy(slice1, slice2)) // hic zat copy edenok len(slice1) == 0
	fmt.Printf("Type %T value %#v\n", slice1, slice1)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(slice1), cap(slice1))

	slice1 = make([]string, 2, 3)
	fmt.Println("Copied", copy(slice1, slice2))
	fmt.Printf("Type %T value %#v\n", slice1, slice1)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(slice1), cap(slice1))

	slice1 = make([]string, len(slice2))
	fmt.Println("Copied: ", copy(slice1, slice2))
	fmt.Printf("Type %T value %#v\n", slice1, slice1)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(slice1), cap(slice1))

	var defaultSlice []string
	fmt.Printf("Type %T value %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(defaultSlice), cap(defaultSlice))

	fmt.Println("Copied: ", copy(defaultSlice, slice2))
	fmt.Printf("Type %T value %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(defaultSlice), cap(defaultSlice))

	rightCopy := make([]string, len(slice2))
	fmt.Println("Copied: ", copy(rightCopy, slice2))
	fmt.Printf("Type %T value %#v\n", rightCopy, rightCopy)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(rightCopy), cap(rightCopy))

	rightCopy2 := append(make([]string, 0, len(slice2)), slice2...)
	fmt.Printf("Type %T value %#v\n", rightCopy2, rightCopy2)
	fmt.Printf("Length: %d Capacity:%d\n\n", len(rightCopy2), cap(rightCopy2))
}

func deleteElementT() {
	slice := []int{1, 2, 3, 4, 5, 6}
	i := 2
	fmt.Printf("Type %T value %#v\n", slice, slice)

	//withAppend := append(slice[:i], slice[i+1:]...) // slice i indeksde 2 sany slice bolyar we i indekssiz catyar
	//fmt.Printf("Type %T value %#v\n", withAppend, withAppend)
	//fmt.Println(slice) // [1,2,3,4,5,6,6]

	// dogry append etmek ucin taze peremennaya alman, appendi  slice ozune dakmaly

	slice = append(slice[:i], slice[i+1:]...)
	fmt.Printf("Type %T value %#v\n", slice, slice)
	fmt.Println(slice) // [1,2,3,4,5,6,6]

	slice = []int{1, 2, 3, 4, 5}
	// withCopy := slice[1:4] // 1-nji indexden baslap 4-nji indexe cenli (4-nji indeksi alanok) value alyp taze slice dakyar
	withCopy := slice[:i+copy(slice[i:], slice[i+1:])]
	fmt.Println(withCopy)

}
