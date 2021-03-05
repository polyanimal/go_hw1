package main

import (
	. "github.com/polyanimal/go_hw1/2/calc"
	"log"
	"os"
)


func main() {
	if len(os.Args) != 2{
		log.Fatal("insufficient number of program arguments")
	}

	exp := os.Args[1]
	exp = InfToPosf(exp)
	println(Calc(exp))
}