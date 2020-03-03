package main 

import (
	"fmt"
)

func main() {
	emails  := make(map[string]string)

	emails["dannon"] = "dannonbigham@gmail.com"
	emails["destiny"] = "destinymalone@gmail.com"

	fmt.Println(emails)
}