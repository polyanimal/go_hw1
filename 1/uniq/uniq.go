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
	k := s

	fields := strings.Fields(s)
	for i, pref := range fields {
		if i >= a.FieldsNum{
			break
		}
		k = strings.TrimLeft(k, " ")
		k = strings.TrimPrefix(k, pref)
	}

	if a.CharsNum != 0 && a.CharsNum <= len(k) {
		k = k[a.CharsNum:]
	}

	if a.CaseInsensitive {
		k = strings.ToLower(k)
	}

	return k
}

func Uniq(ss []string, a Args) []string {
	output := make([]string, 0)
	count := make([]int, 0)
	res := make([]string, 0)

	prev := ss[0]
	count = append(count, 1)
	output = append(output, ss[0])
	l := 0

	for i, s := range ss {
		if i == 0 {
			continue
		}

		if a.key(s) != a.key(prev) {
			prev = s
			count = append(count, 1)
			output = append(output, s)
			l++
		} else {
			count[l]++
		}
	}

	switch {
	case a.Count:
		for i, s := range output {
			res = append(res, strconv.Itoa(count[i])+" "+s)
		}
	case a.Duplicates:
		for i, s := range output {
			if count[i] > 1 {
				res = append(res, s)
			}
		}
	case a.Uniq:
		for i, s := range output {
			if count[i] == 1 {
				res = append(res, s)
			}
		}
	default:
		res = output
	}

	if len(res) > 0 {
		res[len(res)-1] = strings.TrimSuffix(res[len(res)-1], "\n")
	}

	return res
}
