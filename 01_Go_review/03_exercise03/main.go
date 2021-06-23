package main

import (
	"fmt"
	"strings"
)

type person struct {
	fname   string
	lname   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintf("%s is walking.\n", p.fname)
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

func (t truck) transportationDevice() string {
	return "Transporting Products."
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	return "Transporting Human"
}

type transoprtation interface {
	transportationDevice() string
}

func report(t transoprtation) string {
	return t.transportationDevice()
}

func main() {
	// slice exercise
	fmt.Printf("%[1]s %s %[1]s\n", strings.Repeat("-", 10), "Slice Exercise")
	xi := []int{2, 4, 6, 8, 10}

	for i, num := range xi {
		fmt.Printf("#%-4d -> %-4d\n", i, num)
	}
	fmt.Println()

	// map exercise
	fmt.Printf("%[1]s %s %[1]s\n", strings.Repeat("-", 10), "Map Exercise")
	m := map[string]int{
		"John":  15,
		"Kevin": 13,
	}

	for k, v := range m {
		fmt.Printf("key: %-8s value: %d\n", k, v)
	}
	fmt.Println()

	// struce exercise
	fmt.Printf("%[1]s %s %[1]s\n", strings.Repeat("-", 10), "Struct Exercise")
	p1 := person{
		fname:   "Paul",
		lname:   "Tu",
		favFood: []string{"Fish", "Apple"},
	}
	fmt.Println(p1.fname)
	fmt.Println()

	fmt.Println(p1.favFood)

	// method exercise
	fmt.Printf("%[1]s %s %[1]s\n", strings.Repeat("-", 10), "Method Exercise")
	fmt.Println()
	s := p1.walk()
	fmt.Println(s)

	// struct exercise 2
	var (
		t1 truck
		s1 sedan
	)

	t1 = truck{vehicle{2, "white"}, false}
	s1 = sedan{vehicle{4, "red"}, true}

	fmt.Printf("truck: %v\n", t1)
	fmt.Printf("sedan: %v\n", s1)
	s1.transportationDevice()
	t1.transportationDevice()
	report(t1)
	report(s1)

}
