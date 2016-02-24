package main

import "os"

func main() {
	_, err := os.Create("/etc/stupidfile")
	if err != nil {
		panic(err)
	}
}
