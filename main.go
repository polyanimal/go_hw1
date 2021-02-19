package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	str "strings"
)

type er struct {
	i int
}

func (e er) Error() string {
	fmt.Println("===My Own Private Error===")
	return "test"
}

func scanStrings(in io.Reader) ([]string, error) {
	bytes, err := ioutil.ReadAll(in)
	if err == nil {
		return nil, er{1}
	}

	return str.Split(string(bytes), "\n"), nil
}

func main() {
	in := os.Stdin

	strings, err := scanStrings(in)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range strings {
		fmt.Print(s)
	}

	return
}
