package main

import (
	"io"
	"os"
)

func t(out io.Writer) {
	io.WriteString(out, "s")
}

func main() {

	//inputFile := "data.txt"
	//outputFile := "лул.txt"

	//in, _  := os.Open(inputFile)
	out, _ := os.Create("go.txt")

	t(out)

	//scanner := bufio.NewScanner(in)
	////writer := bufio.NewWriter(out)
	//for scanner.Scan() {
	//	s := scanner.Text()
	//	_, err :=  out.WriteString(s)
	//	if err != nil{
	//		fmt.Println(err)
	//	}
	//}
}
