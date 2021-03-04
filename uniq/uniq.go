package uniq

import (
	"strconv"
	"strings"
)

type Args struct {
	Count, Duplicates, Uniq, CaseInsensitive bool
	FieldsNum, CharsNum                      int
}

func (a *Args) key(s string) string {
	k := ""

	fields := strings.Fields(s)
	if a.FieldsNum <= len(fields) {
		k = strings.Join(fields[a.FieldsNum:], "")
	}

	if a.CharsNum <= len(k) {
		k = k[a.CharsNum:]
	}

	if a.CaseInsensitive {
		k = strings.ToLower(k)
	}

	return k
}

func Uniq(ss []string, a Args) []string {
	count := make([]int, 0)
	output := make([]string, 0)
	res := make([]string, 0)

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

	if a.Count {
		for i, s := range output {
			res = append(res, strconv.Itoa(count[i])+" "+s)
		}
	} else if a.Duplicates {
		for i, s := range output {
			if count[i] > 1 {
				res = append(res, s)
			}
		}
	} else if a.Uniq {
		for i, s := range output {
			if count[i] == 1 {
				res = append(res, s)
			}
		}
	} else {
		res = output
	}

	if len(res) > 0 {
		res[len(res)-1] = strings.TrimSuffix(res[len(res)-1], "\n")
	}

	return res
}
