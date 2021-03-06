package calc

import (
	"log"
	"strconv"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *Stack) pop() interface{} {
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return r
}

func (s *Stack) push(c interface{}) {
	s.data = append(s.data, c)
}

func isNumber(c rune) bool {
	return int(c)-'0' >= 0 && int(c)-'0' <= 9
}

func toInt(c rune) int {
	return int(c) - '0'
}

func priority(c string) int {
	if c == "+" || c == "-" {
		return 1
	} else if c == "*" || c == "/" {
		return 2
	}

	return -1
}

func InfToPosf(exp string) string {
	s := Stack{}
	res := ""
	numBuf := ""

	flush := func(res *string, buf *string) {
		if *buf != "" {
			*res += *buf + " "
			*buf = ""
		}
	}

	for _, ch := range exp {
		c := string(ch)
		switch {
		case isNumber(ch):
			numBuf += c
		case ch == '(':
			flush(&res, &numBuf)
			s.push(c)
		case ch == ')':
			flush(&res, &numBuf)
			for s.peek() != string('(') {
				res += s.pop().(string) + " "
			}
			s.pop()
		case ch == ' ':
			flush(&res, &numBuf)
		default:
			flush(&res, &numBuf)
			if len(s.data) == 0 {
				s.push(c)
			} else {
				if priority(s.peek().(string)) >= priority(c) {
					res += s.pop().(string) + " "
					s.push(c)
				} else {
					s.push(c)
				}
			}
		}
	}
	flush(&res, &numBuf)

	for len(s.data) != 0 {
		res += s.pop().(string) + " "
	}

	return res
}

func eval(op string, s *Stack) int {
	switch op {
	case "+":
		return s.pop().(int) + s.pop().(int)
	case "*":
		return s.pop().(int) * s.pop().(int)
	case "-":
		t := s.pop().(int)
		return s.pop().(int) - t
	case "/":
		t := s.pop().(int)
		return s.pop().(int) / t
	default:
		log.Fatal("Wrong argument for eval")
	}

	return -1
}

func Calc(exp string) int {
	s := Stack{}
	numBuf := ""

	pushInt := func (s *Stack, buf *string) {
		if *buf != "" {
			i, err := strconv.Atoi(numBuf)
			if err != nil {
				log.Fatal(err)
			}
			s.push(i)

			*buf = ""
		}
	}

	for _, ch := range exp {
		switch {
		case isNumber(ch):
			numBuf += string(ch)
		case ch == ' ':
			pushInt(&s, &numBuf)
		default:
			val := eval(string(ch), &s)
			s.push(val)
		}

	}

	return s.pop().(int)
}
