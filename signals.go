package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		fmt.Println("Recived siganl: ", sig)
		done <- true
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stop work due to signal.")
				return
			default:
				fmt.Println("Working")
				time.Sleep(time.Second)
			}
		}

		// // sig := <-sigs
		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Recived SIGINT (Interrupt)")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Recived SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Recived SIGHUP (Hang Up)")
		// 	case syscall.SIGUSR1:
		// 		fmt.Println("Recived SIGUSR1 (User Defined Signal 1)")
		// 		fmt.Println("User defined function is executed")
		// 		continue
		// 	}
		// 	fmt.Println("Grafeful exit.")

		// 	os.Exit(0)
		// }

	}()

	fmt.Println("Working..., pid: ", pid)
	for {
		time.Sleep(time.Second)
	}
}
