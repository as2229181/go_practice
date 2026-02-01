package main

import "fmt"
type Person struct{
	firstName string
	lastName string
	email string
	age int
}
func (p Person) isAdult() bool {
	if p.age >= 18{
		return true
	}
	return false
}

func (p *Person) addAgeByOne() {
	p.age ++
}

func main(){
	p := Person{
		firstName: "Ji-Shin",
		lastName: "Kuo",
		email: "as2229181@gmail.com",
		age: 18,
	}
	s := fmt.Sprint(p.lastName, " ",p.firstName, "is", p.age, "years old")
	fmt.Println(s)
	fmt.Println(p.isAdult())
	fmt.Println(p.age)
	p.addAgeByOne()
	fmt.Println(p.age)
}