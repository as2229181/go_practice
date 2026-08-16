package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     By
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	age := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sort Result", people)

	// ======== Sort.Slice
	stringSlice := []string{"banana", "apple", "cherry"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last charactter:", stringSlice)

	// sort.Sort(ByAge(people))

	// numbers := []int{5, 4, 3, 1, 2}
	// // sort.Ints(numbers)
	// // sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	// slices.Sort(numbers)
	// fmt.Println("Sorted numbers", numbers)
	// names := []string{"john", "anthony", "steve", "victor", "walter"}
	// sort.Strings(names)
	// fmt.Println("Sorted strings", names)
}
