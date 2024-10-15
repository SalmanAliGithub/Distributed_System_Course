package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 20.5
	var c string = "Golang"
	var d bool = true

	e := "Inferred declaration"

	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)
	fmt.Println("Inferred Declaration:", e)

}
