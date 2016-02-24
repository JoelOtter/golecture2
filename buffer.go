package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "First message"
	messages <- "Second message"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
