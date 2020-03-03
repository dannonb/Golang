package main 

import (
	"fmt"
)

func main() {
	a := 5 
	b := &a 

	fmt.Println(a)
	fmt.Println(*b)

	a = 6
	fmt.Println(a)
	fmt.Println(*b)
}