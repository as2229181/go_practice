package main

import (
	"fmt"
	"reflect"
)

type Greeter struct{}

func (g Greeter) Greet(name string) string {
	return "Hello " + name
}

func main() {
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)

	var method reflect.Method

	fmt.Println("Type:", t)
	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}
	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice")})

	fmt.Println("Greet result", results[0].String())
}

// type person struct {
// 	Name string
// 	age  int
// }

// func main() {
// 	p := person{Name: "Alice", age: 30}
// 	v := reflect.ValueOf(p)

// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}

// 	v1 := reflect.ValueOf(&p).Elem()
// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v1.Field(i))
// 	}

// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 		fmt.Println(nameField)
// 	} else {
// 		fmt.Println("Cannot Set")
// 	}
// }

// func main() {
// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())

// 	y := 10
// 	a1 := reflect.ValueOf(&y)
// 	fmt.Println(a1)
// 	fmt.Println("type ", a1.Type())
// 	v1 := reflect.ValueOf(&y).Elem()
// 	fmt.Println(v1)
// 	fmt.Println("type ", v1.Type())
// 	v1.SetInt(18)
// 	fmt.Println(v1)

// 	var itf interface{}
// 	itf = "Hello"
// 	v3 := reflect.ValueOf(itf)
// 	fmt.Println("V3 type", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String value", v3.String())
// 	}

// }
