package main

import "fmt"

func printTenTimes(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name)
	}
}

func main() {
	go printTenTimes("Alice")
	go printTenTimes("Bob")
	go printTenTimes("Carluccio")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
