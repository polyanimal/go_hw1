package main

import (
	"bufio"
	"flag"
	. "github.com/polyanimal/go_hw1/1/uniq"
	"io"
	"log"
	"os"
)


func main() {
	cFlag := flag.Bool("c", false, "Count occurrences")
	dFlag := flag.Bool("d", false, "Duplicated strings")
	uFlag := flag.Bool("u", false, "Unique strings")
	iFlag := flag.Bool("i", false, "Case insensitive")
	fNum := flag.Int("f", 0, "Skip first N words")
	sNum := flag.Int("s", 0, "Skip first N symbols")

	flag.Parse()
	if *cFlag && *dFlag || *cFlag && *uFlag || *dFlag && *uFlag {
		log.Fatal("\nFlags -c, -d and -u are exclusive and can't be used together\n")
	}
	if *fNum < 0 {
		log.Fatalf("\n%d : invalid number of fields to skip\n", *fNum)
	}
	if *sNum < 0 {
		log.Fatalf("\n%d : invalid number of bytes to skip\n", *sNum)
	}

	in := os.Stdin
	out := os.Stdout

	inputFile := ""
	outputFile := ""
	if len(flag.Args()) > 0 {
		inputFile = flag.Args()[0]
	}
	if len(flag.Args()) > 1 {
		outputFile = flag.Args()[1]
	}

	var err error = nil
	if inputFile != "" {
		in, err = os.Open(inputFile)

		if err != nil {
			log.Fatal(err)
		}
	}
	if outputFile != "" {
		out, err = os.Create(outputFile)

		if err != nil {
			log.Fatal(err)
		}
	}

	ss := scanStrings(in)
	res := Uniq(ss, Args{
		Count:           *cFlag,
		Duplicates:      *dFlag,
		Uniq:            *uFlag,
		CaseInsensitive: *iFlag,
		FieldsNum:       *fNum,
		CharsNum:        *sNum,
	})

	printSs(res, out)
	return
}

func scanStrings(in io.Reader) []string {
	ss := make([]string, 0)
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		ss = append(ss, scanner.Text()+"\n")
	}

	return ss
}

func printSs(ss []string, out io.Writer) {
	for _, s := range ss {
		io.WriteString(out, s)
	}
}
