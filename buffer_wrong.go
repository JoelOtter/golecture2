package main

import "fmt"

func main() {
	messages := make(chan string)
	messages <- "Hello!"
	fmt.Println(<-messages)
}
