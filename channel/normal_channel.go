package channel

import (
	"fmt"
	"time"
)


func normal_channel() {
	greeting := make(chan string)
	go func(){
		greeting <- "string"
		greeting <- "string 2"
		close(greeting)
	}()

	// go func(){
	// 	reciver := <- greeting
	// 	fmt.Println(reciver)
	// 	reciver = <- greeting
	// 	fmt.Println(reciver)
	// 	}()

	for msg := range greeting {
		fmt.Println("接收到:", msg)
	}


	time.Sleep(2 * time.Second)
	fmt.Println("end of main")
	}
	