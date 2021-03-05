package calc

import (
	"log"
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

	for _, ch := range exp {
		c := string(ch)
		if isNumber(ch) {
			res += c
		} else if ch == '(' {
			s.push(c)
		} else if ch == ')' {
			for s.peek() != string('(') {
				res += s.pop().(string)
			}
			s.pop()
		} else { //  + - * /
			if len(s.data) == 0 {
				s.push(c)
			} else {
				if priority(s.peek().(string)) >= priority(c) {
					res += s.pop().(string)
					s.push(c)
				} else {
					s.push(c)
				}
			}
		}
	}

	for len(s.data) != 0 {
		res += s.pop().(string)
	}

	return res
}

func eval(op string, s * Stack) int {
	switch op {
	case "+":
		a := s.pop().(int)
		b := s.pop().(int)
		return a + b
	case "*":
		a := s.pop().(int)
		b := s.pop().(int)
		return a * b
	case "-":
		a := s.pop().(int)
		b := s.pop().(int)
		return b - a
	case "/":
		a := s.pop().(int)
		b := s.pop().(int)
		return b / a
	default:
		log.Fatal("Wrong argument for eval")
	}

	return -1
}

func Calc(exp string) int {
	s := Stack{}

	for _, ch := range exp {
		ss := string(ch)
		if isNumber(ch) {
			s.push(toInt(ch))
		} else {
			val := eval(string(ch), &s)
			s.push(val)
		}

		ss += "2"
	}

	return s.pop().(int)
}
