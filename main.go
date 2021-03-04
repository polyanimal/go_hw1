package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	count, duplicates, uniq, caseInsensitive bool
	fieldsNum, charsNum                      int
}

func (a *Args) key(s string) string {
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

func uniq(ss []string, a Args) []string {
	count  := make([]int,    0)
	output := make([]string, 0)
	res    := make([]string, 0)

	prev := ""
	l := -1

	for _, s := range ss {
		if a.key(s) != a.key(prev) {
			prev = s
			count = append(count, 1)
			output = append(output, s)
			l++
		} else {
			count[l]++
		}
	}

	if a.count {
		for i, s := range output {
			res = append(res, strconv.Itoa(count[i]) + " " + s)
		}
	} else if a.duplicates {
		for i, s := range output {
			if count[i] > 1 {
				res = append(res, s)
			}
		}
	} else if a.uniq {
		for i, s := range output {
			if count[i] == 1 {
				res = append(res, s)
			}
		}
	} else {
		res = output
	}

	if len(res) > 0{
		res[len(res) - 1] = strings.TrimSuffix(res[len(res) - 1], "\n")
	}

	return res
}

func main() {
	cFlag := flag.Bool("c", false, "Count occurrences")
	dFlag := flag.Bool("d", false, "Duplicated strings")
	uFlag := flag.Bool("u", false, "Unique strings")
	iFlag := flag.Bool("i", false, "Case insensitive")
	fNum  := flag.Int("f", 0, "Skip first N words")
	sNum  := flag.Int("s", 0, "Skip first N symbols")

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
	res := uniq(ss, Args{
		count:           *cFlag,
		duplicates:      *dFlag,
		uniq:            *uFlag,
		caseInsensitive: *iFlag,
		fieldsNum:       *fNum,
		charsNum:        *sNum,
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
	//for _, i := range c {
	//	io.WriteString(out, strconv.Itoa(i))
	//}
	for _, s := range ss {
		io.WriteString(out,s)
	}
}
