package main

import (
	"fmt"

	"github.com/ClaudioRoncaglio/puppy"
)

func main() {
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())

	// anche così
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

}
