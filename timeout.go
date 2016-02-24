package main

import "fmt"
import "time"

func sleepThenTalk(seconds int, say string, c chan<- string) {
	time.Sleep(time.Second * time.Duration(seconds))
	c <- say
}

func main() {
	c := make(chan string)

	go sleepThenTalk(50, "Big wait!", c)

waitloop:
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-time.After(time.Second * 3):
			fmt.Println("I'm bored...")
			break waitloop
		}
	}
}
