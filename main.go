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
	count, duplicates, uniq, caseInsensitive bool
	fieldsNum, charsNum                      int
	inputFile, outputFile                    string
}

func (a * Args) key(s string) string {
	k := ""

	fields := strings.Fields(s)
	if a.fieldsNum <= len(fields) {
		k = strings.Join(fields[a.fieldsNum:], "")
	}

	if a.charsNum <= len(k) {
		k = k[a.charsNum:]
	}

	if a.caseInsensitive {
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

	//cFlag := flag.Bool("-count", false, "Count occurrences")
	//dFlag := flag.Bool("-duplicates", false, "Duplicated strings")
	//uFlag := flag.Bool("-count", false, "Unique strings")
	//iFlag := flag.Bool("-caseInsensitive", false, "Case insensitive")

	//fFlag := flag.Int("-fieldsNum", 0, "Skip first N words")
	//sFlag := flag.Int("-charsNum", 0, "Skip first N symbols")

	//inputFile := flag.String("input_file", "", "caseInsensitive")
	//outputFile := flag.String("output_file", "", "o")

	strs, err := scanStrings(in)
	if err != nil {
		log.Fatal(err)
	}

	res, _ := uniq(strs, Args{
		count:           false,
		duplicates:      false,
		uniq:            false,
		caseInsensitive: false,
		fieldsNum:       0,
		charsNum:        0,
		inputFile:       "",
		outputFile:      "",
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