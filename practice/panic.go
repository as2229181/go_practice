package main

import (
	"fmt"
)

func main(){
	process(10)
	process(-3)
}

func process(input int){
	defer fmt.Println("1")
	defer fmt.Println("2")
	if input < 0 {
		fmt.Println("before panic")
		panic("input is less than zero")
	}
	fmt.Println("input", input)
}