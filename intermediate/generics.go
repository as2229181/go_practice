package main

import "fmt"

// func swap[T any](x, y T) (T, T) {
// 	return y, x
// }

// func main() {
// 	x, y := 1, 2
// 	new_x, new_y := swap(x, y)
// 	fmt.Println(new_x, new_y)
// }

// type Stack[T any] struct {
// 	element []T
// }

// func (s *Stack[T]) push(element T) {
// 	s.element = append(s.element, element)
// }

// func (s *Stack[T]) pop() (T, bool){
// 	if len(s.element) == 0 {
// 		var zero T
// 		return zero, false
// 	}
// 	element := s.element[len(s.element)-1]
// 	s.element = s.element[: len(s.element)-1]
// 	return element, true
// }

// func (s Stack[T]) isEmpty() bool {
// 	return len(s.element) == 0
// }

// func (s Stack[T]) printAll() {
// 	if len(s.element) == 0{
// 		fmt.Println()
// 	}
// 	for _, value := range s.element {
// 		fmt.Print(value)
// 	}
// 	fmt.Println()
// }

// func main() {
// 	stack1 := Stack[int]{}
// 	stack1.push(1)
// 	stack1.push(3)
// 	fmt.Println(stack1.isEmpty())
// 	stack1.printAll()
// 	stack1.pop()
// 	stack1.pop()
// 	fmt.Println(stack1.isEmpty())
// }

func modifyValue(ptr *int) {
    *ptr = 20
}

func main() {
    value := 10
    ptr := &value

    fmt.Println("Before modification:", value)

    modifyValue(ptr)

    fmt.Println("After modification:", value)
}