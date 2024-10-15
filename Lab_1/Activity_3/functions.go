package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)

	fmt.Println("Sum:", result)

	if result > 5 {
		fmt.Println("Result is greater that 5")
	} else {
		fmt.Println("Result is 5 or less")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}
}
