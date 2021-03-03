package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func scanStrings(in io.Reader) ([]string, error) {
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	return strings.SplitAfter(string(bytes), "\n"), nil
}

func key(s string, f int, c int, i bool ) string {
	k := strings.Join(strings.Fields(s)[f:], "")
	k = k[c:]
	if i {
		k = strings.ToLower(k)
	}

	return k
}

func uniq(strs []string) ([]string, map[string]int) {
	output := make([]string, 0)
	m := make(map[string]int)
	prev := ""

	for _, s := range strs {
		m[key(s, 1, 0, false)]++
		if key(s, 1, 0, false) != key(prev, 1, 0, false) {
			prev = s
			output = append(output, s)
		}
	}

	return output, m
}

func print(strs []string) {
	fmt.Println("----------------")
	for _, s := range strs {
		fmt.Print(s)
	}
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

	res, _ := uniq(strs)


	fmt.Println(res)
	return
}
