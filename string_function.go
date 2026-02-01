package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	num := 19
	// 	str_1 := strconv.Itoa(num)
	// 	fmt.Println(reflect.TypeOf(str_1))
	// 	fmt.Println(len(str_1))

	// 	fruits := "banana, apple, orange"
	// 	parts :=  strings.Split(fruits, ", ")
	// 	fmt.Println(parts)
	// 	countries := []string{"Germany", "America", "Brazil"}
	// 	joinedCountries := strings.Join(countries, ", ")
	// 	fmt.Println(joinedCountries)

	// 	str := "hello world"
	// 	fmt.Println(strings.Contains(str, "Go"))

	// 	replacedStr := strings.Replace(str, "hello", "Go", 1)
	// 	fmt.Println(replacedStr)
	// 	fmt.Println(strings.Repeat("xyz", 3))

	// 	fmt.Println(strings.Count("hello", "l"))
	// 	fmt.Println(strings.HasPrefix("hello_world", "hello"))

	// 	str2 := "hello world 123123 yeah"
	// 	re := regexp.MustCompile(`\d+`)
	// 	matches := re.FindAllString(str2, -1)
	// 	fmt.Println(matches)

	// 	str3 := "hello 世界"
	// 	fmt.Println(utf8.RuneCountInString(str3))
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString(",")
	builder.WriteString("world")
	result:= builder.String()
	fmt.Println(result)
	builder.Reset()
	fmt.Println(result)
}