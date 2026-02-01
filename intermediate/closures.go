package main

import "fmt"

func main() {
	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// sequence2 := adder()
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	subTracer := func() func()int{
		countDown := 99
		return func() int {
			countDown -= 1
			return countDown
		}
	}()
	fmt.Println(subTracer())
}
func adder() func() int {
	i := 0
	fmt.Println("Previous value of i is ", i)
	return func() int {
		i++
		fmt.Println("Now value of i is ", i)
		return i
	}
}