package main

import "fmt"

func main() {
	a := 5
	b := &a

	// b printed out memory address of a
	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)

	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
