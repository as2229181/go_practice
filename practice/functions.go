package practice

import (
	"fmt"
)

func main() {
	// sum := add(1, 2)
	// fmt.Println(sum)
	// greet := func(){
	// 	fmt.Println("hello in anonymous function")
	// }
	// greet()
	addResult := applyOperation(3, 4, add)
	fmt.Println(addResult)
	multiplier2 := applyMultiplier(2)
	result := multiplier2(6)
	fmt.Println(result)
}
func add(a, b int) int {
	return a + b
}

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}
func applyMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}