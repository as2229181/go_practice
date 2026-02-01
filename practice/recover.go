package main

import "fmt"
func main() {
	process()
}
func process(){
	defer func(){
		r := recover()
		fmt.Println(r)
		if r != nil{
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("start process")
	// panic("Something went wrong")

}