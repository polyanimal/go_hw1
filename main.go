package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Args struct {
	c, d, u, i bool
	f, s int
	inputFile, outputFile string
}

func (a * Args) key(s string) string {
	k := ""

	fields := strings.Fields(s)
	if a.f <= len(fields) {
		k = strings.Join(fields[a.f:], "")
	}

	if a.s <= len(k) {
		k = k[a.s:]
	}

	if a.i {
		k = strings.ToLower(k)
	}

	return k
}

func uniq(strs []string, a Args) ([]string, map[string]int) {
	output := make([]string, 0)
	m := make(map[string]int)
	prev := ""

	for _, s := range strs {
		m[a.key(s)]++
		if a.key(s) != a.key(prev) {
			prev = s
			output = append(output, s)
		}
	}

	return output, m
}

func main() {
	in := os.Stdin

	//cFlag := flag.Bool("-c", false, "Count occurrences")
	//dFlag := flag.Bool("-d", false, "Duplicated strings")
	//uFlag := flag.Bool("-c", false, "Unique strings")
	//iFlag := flag.Bool("-i", false, "Case insensitive")

	//fFlag := flag.Int("-f", 0, "Skip first N words")
	//sFlag := flag.Int("-s", 0, "Skip first N symbols")

	//inputFile := flag.String("input_file", "", "i")
	//outputFile := flag.String("output_file", "", "o")

	strs, err := scanStrings(in)
	if err != nil {
		log.Fatal(err)
	}

	res, _ := uniq(strs, Args{
		c:          false,
		d:          false,
		u:          false,
		i:          false,
		f:          0,
		s:          0,
		inputFile:  "",
		outputFile: "",
	})


	fmt.Println(res)
	return
}


func scanStrings(in io.Reader) ([]string, error) {
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	return strings.SplitAfter(string(bytes), "\n"), nil
}

func printSs(strs []string) {
	fmt.Println("----------------")
	for _, s := range strs {
		fmt.Print(s)
	}
}