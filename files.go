package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// create file
	//ioutilPackage()
	//osPackage()

	// delete file
	removeFile()

}

func ioutilPackage() {
	//Create new file
	data := []byte("Some text .... ")
	err := ioutil.WriteFile("text1.txt", data, 0600)

	if err != nil {
		fmt.Println("File cannot create\n", err)
	}
	file_data, ok := ioutil.ReadFile("text1.txt")
	if ok != nil {
		fmt.Println("Error")
	}
	//fmt.Println(file_data)
	fmt.Println(string(file_data))
}

func osPackage() {
	file, err := os.Create("textFile.txt")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer file.Close()
	data := "Hello\nWorld\nfrom\n\tFile"
	file.WriteString(data)

	file_data, _ := os.ReadFile(file.Name())
	fmt.Println(string(file_data))
}

func removeFile() {
	os.Remove("text1.txt")
}
