package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, Paul"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `Shaken, not stirred.`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// slice
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	// map
	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	// struct
	p1 := person{
		"Miss", "Paula",
	}
	fmt.Println(p1)

	// function
	p1.speak()

	// composition
	sa1 := secretAgent{
		person{
			fname: "James",
			lname: "Bond",
		},
		true,
	}
	sa1.speak()

	// interface
	saySomething(sa1)
	saySomething(p1)
}
