package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func ticketProcess(requests <-chan ticketRequest, result chan<- int) {
	for request := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with cost %d\n", request.numTickets, request.personID, request.cost)
		time.Sleep(time.Second)
		result <- request.personID
	}

}

func main() {
	numRequest := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequest)
	ticketResult := make(chan int)
	for range 3 {
		go ticketProcess(ticketRequests, ticketResult)
	}

	for i := range numRequest {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * 2 * price}
	}
	close(ticketRequests)

	for range numRequest {
		fmt.Printf("Ticket for person Id %d process successfully!\n", <-ticketResult)
	}
}

//  ===========================Baisc worker pool
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processingt task %d\n", id, task)
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {
// 	numWokers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for i := range numWokers {
// 		go worker(i, tasks, results)
// 	}

// 	for i := range numJobs {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}

// 	close(results)
// 	fmt.Println("End of program")
// }
