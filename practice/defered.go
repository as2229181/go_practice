package main

import "fmt"

func main() {
	numbers := []int{1 , 2, 3, 4}
	// process(10)
	calculate(numbers ...)
}
func calculate(num ...int){
	total := 0
	for _, value := range num{
		total += value
	}
	fmt.Println(total)
}
// func process(i int){
// 	defer fmt.Println("i", i)
// 	defer fmt.Println("defer 1")
// 	defer fmt.Println("defer 2")
// 	defer fmt.Println("defer 3")
// 	i ++
// 	fmt.Println("value: ", i)
// }
