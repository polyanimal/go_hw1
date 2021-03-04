package main

import "fmt"

func main() {
	count := make([]int, 0)
	l := 0

	for i :=0; i < 10; i++{
		count = append(count, 1)
		count[l]++
		l++
	}

	fmt.Print(count)
}
