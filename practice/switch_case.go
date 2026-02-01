package practice

import (
	"fmt"
)

func main() {
	// fruit := "orange"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("It is an apple")
	// case "banana":
	// 	fmt.Println("It is a banana")
	// default:
	// 	fmt.Println("Unknown")
	// }
	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is bigger than 20")
	// }
	checkType(20)
	checkType("yes")
}

func checkType(x interface{}) {
	switch x.(type){
	case int:
		fmt.Println("it is integer")
	case float64:
		fmt.Println("It is float 64")
	case string:
		fmt.Println("It is  string")
	default:
		fmt.Println("Unknown type")
	}
}