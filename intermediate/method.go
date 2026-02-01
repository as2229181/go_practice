package main

import "fmt"

func main() {

	num :=	MyInt(1)
	num1 := MyInt(-2)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.WelcomeMessage())
}

type MyInt int

func (m MyInt) IsPositive() bool{
	return m > 0
}

func (MyInt) WelcomeMessage() string {
	return "welcome to My Int Type"
}