package main

import "fmt"

func main() {

	var a string = "42"
	var b int = 42
	var c float64 = 42.00

	fmt.Printf("Variabile a di tipo %T vale %v\n", a, a)
	fmt.Printf("Variabile b di tipo %T vale %v\n", b, b)
	fmt.Printf("Variabile c di tipo %T vale %v\n", c, c)
}
