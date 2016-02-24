package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Hello!"
	}()

	msg := <-messages
	fmt.Println(msg)
}
