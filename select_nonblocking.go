package main

import "fmt"
import "time"

func sleepThenTalk(seconds int, say string, c chan<- string) {
	time.Sleep(time.Second * time.Duration(seconds))
	c <- say
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go sleepThenTalk(5, "Five!", c2)
	go sleepThenTalk(2, "Two!", c1)

	received := 0
	for received < 2 {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
			received++
		case msg2 := <-c2:
			fmt.Println(msg2)
			received++
		default:
			fmt.Println("Waiting...")
		}
	}
}
