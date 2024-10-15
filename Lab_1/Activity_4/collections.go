package main

import "fmt"

func main() {
	// Array
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice
	slice := []int{4, 5, 6}
	fmt.Println("Slice:", slice)
	slice = append(slice, 7)
	fmt.Println("Appended slice:", slice)

	// Map
	myMap := make(map[string]int)
	myMap["Alice"] = 25
	myMap["Bob"] = 30
	fmt.Println("Map:", myMap)
	fmt.Println("Alice's age:", myMap["Alice"])

	// Slice looping
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d", index, value)
	}

	// Map looping
	for key, value := range myMap {
		fmt.Printf("%s's age is %d\n", key, value)
	}

}
