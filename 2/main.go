package main

import (
	"fmt"
	"github.com/polyanimal/go_hw1/2/calc"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("insufficient number of program arguments")
	}

	exp, err := calc.InfToPosf(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calc.Calc(exp))
}
