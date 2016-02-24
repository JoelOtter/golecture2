package main

import (
	"fmt"
	"os"
)

func createFile(name string) *os.File {
	fmt.Printf("Creating file with path %v\n", name)
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File, data string) {
	fmt.Printf("Writing data [%v] to file\n", data)
	fmt.Fprintln(file, data)
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func main() {
	f := createFile("./defer_is_fun.txt") // f is a pointer
	defer closeFile(f)
	writeFile(f, "Defer is fun!")
}
