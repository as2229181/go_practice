package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	todoContext := context.TODO()
// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	contextBkg := context.Background()
// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))

// }
// func checkEventOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}

// }

// func contextCancel(cl context.CancelFunc) {
// 	cl()
// 	fmt.Println("cancel context")

// }

// func main() {
// 	ctx := context.TODO()

// 	result := checkEventOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO():", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 3*time.Millisecond)
// 	// time.Sleep(3 * time.Millisecond)
// 	result = checkEventOdd(ctx, 10)
// 	fmt.Println("Result from timeout context", result)
// 	// time.Sleep(1 * time.Second)
// 	result = checkEventOdd(ctx, 11)
// 	fmt.Println("Result after timeout context timeout", result)
// 	defer contextCancel(cancel)
// }

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled", ctx.Err())
			return
		default:
			fmt.Println("Working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	ctx = context.WithValue(ctx, "requestId", "ashjkfhasjkldhkajl")
	go doWork(ctx)
	time.Sleep(5 * time.Second)
	requestId := ctx.Value("requestId")
	if requestId != nil {
		fmt.Println("Request id:", requestId)
	} else {
		fmt.Println("No request ID fuound")
	}
}
