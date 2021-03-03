package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1"

	z := strings.Fields(s)

	//for i := range z {
	//	fmt.Printf("%s\n", z[i])
	//}

	fmt.Println(len(z))
}