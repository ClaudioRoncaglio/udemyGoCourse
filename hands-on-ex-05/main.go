package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 42
	fmt.Println(z)

	a, b := 43, "Happiness"
	fmt.Println(a, b)

	var c float32 = 42.42
	fmt.Printf("%v is of type %T\n", c, c)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)

}
