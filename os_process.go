package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "5")

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	time.Sleep(2 * time.Second)
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error waiting:", err)
		return
	}

	fmt.Println("Process is killed")

	// waiting
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("Error waiting:", err)
	// 	return
	// }

	// fmt.Println("Process is complete")
	// cmd := exec.Command("grep", "foo")

	// cmd.Stdin = strings.NewReader("foo is result\nbar\nbaz\n")

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Output: ", string(output))
	// cmd := exec.Command("echo", "Hello World !!")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// fmt.Println("Output: ", string(output))
}
