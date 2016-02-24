package main

import "fmt"

func printTenTimes(name string, done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(name)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go printTenTimes("Alice", done)
	<-done
}
