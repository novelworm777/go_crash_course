package main

import "fmt"

func main() {
	ids := []int{33, 73, 13, 17, 3, 7}

	// Loops through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharon": "sharon@gmail.com",
	}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Println("Name: " + key)
	}
}
