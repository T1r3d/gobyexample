package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "start job", job)
		time.Sleep(time.Second)
		results <- job * 2
		fmt.Println("worker", id, "finished job", job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= 5; r++ {
		<-results
	}
}
