package main

import "fmt"

type person struct {
	name   string
	age    int
	height int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 25
	p.height = 175
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20, 200})

	fmt.Println(person{name: "Alice", age: 30, height: 155})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40, height: 165})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50, height: 140}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
