package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker ID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("Worker ID %d finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}
	wg.Wait()
	fmt.Println("finished")
}

// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker ID %d starting\n", id)
// 	time.Sleep(time.Second)
// 	for task := range tasks {
// 		results <- task * 2
// 	}
// 	fmt.Printf("Worker ID %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 6
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	go func() {
// 		wg.Wait()      // need this to wait all worker are done
// 		close(results) // otherwise results might close before work done
// 	}()

// 	for result := range results {
// 		fmt.Println("Results:", result)
// 	}

// 	fmt.Println("Program Done")
// }

//  ==================== Basic example
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d start\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All worker done\n")
// }
