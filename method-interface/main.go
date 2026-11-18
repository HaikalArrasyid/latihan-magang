package main

import (
	"fmt"
	"math"
)

// method
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// interface
type NamaSaya interface {
	GetName() string
}

func sayHello(namaSaya NamaSaya) {
	fmt.Println("hello", namaSaya.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {

	//method
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//pointer
	name := "bagas"
	name2 := &name
	*name2 = "agus"

	println(*name2)

	//interface
	var agus Person
	agus.Name = "agus saputra"

	sayHello(agus)

}

//testing git fetch