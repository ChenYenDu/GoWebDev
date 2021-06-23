package main

import (
	"fmt"
	"math"
	"strings"
)

//  hands on 1
type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

// hands on 2
type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println("Hahaha i am a person")
}

func (p person) speak() string {
	return fmt.Sprintln("Hahaha i am a human")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("Haha i am a secret agent")
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln("Haha i am a human")
}

type human interface {
	speak() string
}

func vomit(h human) {
	fmt.Printf("%T\n", h)
	fmt.Println(h)
	switch v := h.(type) {
	case person:
		fmt.Println(v.fname)
	case secretAgent:
		fmt.Println(v.fname)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	// hands on 1
	fmt.Println("Hands on 1")
	sq := square{20}
	cl := circle{4}

	info(sq)
	info(cl)

	// hands on 2
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println("Hands on 2")
	p1 := person{"Paul", "Tu"}
	sa1 := secretAgent{person{"James", "Bond"}, true}

	p1.pSpeak()
	sa1.saSpeak()

	// hands on 3
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println("Hands on 3")
	vomit(p1)
	vomit(sa1)

}
