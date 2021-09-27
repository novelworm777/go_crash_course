package main

import "fmt"

// func function_name(parameters) return_type {}
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Mika"))
	fmt.Println(getSum(3, 4))
}
