package practice

import "fmt"

func main() {
	var number [5]int
	fmt.Println(number)
	number[4] = 20
	fmt.Println(number)
	fruit := [4]string{"banana", "orange"}
	fmt.Println(fruit)
	var list = [5]int{1, 2, 3, 4, 5}
	fmt.Println(list)
	for i := 0; i < len(number); i++{
		fmt.Println("Element at index", i, ":", number[i])
	}
	for index, value := range number {
		fmt.Printf("Index : %d Value : %d \n", index, value)
	}
}