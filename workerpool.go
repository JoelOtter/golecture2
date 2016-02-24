package main

import "fmt"
import "time"

const numJobs int = 90000
const numWorkers int = 90000

func squarer(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %v processing job %v\n", id, j)
		time.Sleep(time.Second)
		results <- j * j
	}
}

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go squarer(w, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for r := 1; r <= numJobs; r++ {
		fmt.Printf("Got result: %v\n", <-results)
	}
}
