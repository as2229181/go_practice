package stateful_goroutine

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	count int
	ch    chan int
}

func (w *StatefulWorker) start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)
			}
		}

	}()
}

func (w *StatefulWorker) send(value int) {
	w.ch <- value
}

func main() {
	statefulWorker := &StatefulWorker{
		ch: make(chan int),
	}
	statefulWorker.start()
	for i := range 5 {
		statefulWorker.send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
