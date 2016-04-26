package main

import "fmt"

type T struct {
	name  string // name of object
	value int    // its value
}

func sayHi(name string) {
	fmt.Printf("Hello, %v\n", name)
}

func main() {
	t := new(T)
	t.name = "joe"
	t.value = 1

	fmt.Printf("Example 01\n")
	fmt.Printf("name: %v, value: %v\n", t.name, t.value)
	fmt.Printf("t -  type: %T, value: %v\n", t, t)
	fmt.Printf("*t - type: %T, value: %v\n", *t, *t)

	sayHi("Joe")
}
