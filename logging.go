package main

import (
	"log"
	"os"
)

func main() {
	// N-1
	//log.Println("Hello from logging print")
	//log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix | log.Llongfile | log.LUTC)
	//log.Println("Hello from logging print")
	//// log.Panic("Panic")
	//log.Fatalln("Fatal Panic") // Panic den sonra Recover ulanmayan bolsak Fatal ulanylyp bilner

	// N-2
	// Consola cykarylyan habary terminala dalde, fayla cykaryp bolyar.
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//file, err := os.Create("info.log")
	//if err != nil {
	//	log.Println("ERROR")
	//}
	//log.SetOutput(file)
	//log.Println("Log in the file")
	//file.Close()
	//
	//log.SetOutput(os.Stdout)
	//log.Println("Print in standard output again")

	// N - 3
	flags := log.LstdFlags | log.Lshortfile

	// create loggers with New
	logInfo := log.New(os.Stdout, "INFO:", log.Ldate)
	logWarn := log.New(os.Stdout, "WARN:", flags)
	logErr := log.New(os.Stdout, "ERR:", log.Llongfile)

}
