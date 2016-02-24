package main

import "fmt"

func main() {
	c := make(chan string, 3)
	c <- "Alice"
	c <- "Bob"
	c <- "Carlos"
	close(c) // What happens without this?

	for name := range c {
		fmt.Println(name)
	}
}
