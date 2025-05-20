package main

import (
	"fmt"

	"github.com/ClaudioRoncaglio/puppy"
)

func main() {
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	// anche così
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

}
