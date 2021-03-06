package calc

import (
	"errors"
	"log"
	"strconv"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) peek() interface{} {
	if len(s.data) == 0 {
		log.Println("stack underflow")
		return nil
	}

	return s.data[len(s.data)-1]
}

func (s *Stack) pop() interface{} {
	if len(s.data) == 0 {
		log.Println("stack underflow")
		return nil
	}

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

func priority(c string) int {
	if c == "+" || c == "-" {
		return 1
	} else if c == "*" || c == "/" {
		return 2
	}

	return -1
}

func checkParenthesis(s string) bool {
	counter := 0
	for _, c := range s {
		switch {
		case counter < 0:
			return false
		case c == '(':
			counter++
		case c == ')':
			counter--
		}
	}

	return counter == 0
}

func InfToPosf(exp string) (string, error) {
	if !checkParenthesis(exp) {
		return "", errors.New("bad syntax: parenthesis don't match")
	}

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

	return res, nil
}

func eval(op string, s *Stack) (int, error) {
	switch op {
	case "+":
		return s.pop().(int) + s.pop().(int), nil
	case "*":
		return s.pop().(int) * s.pop().(int), nil
	case "-":
		t := s.pop().(int)
		return s.pop().(int) - t, nil
	case "/":
		t := s.pop().(int)
		if t == 0 {
			return -1, errors.New("division by zero")
		}
		return s.pop().(int) / t, nil
	default:
		return -1, errors.New("bad syntax: unsupported operation")
	}
}

func Calc(exp string) (int, error) {
	s := Stack{}
	numBuf := ""

	pushInt := func(s *Stack, buf *string) {
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
			val, err := eval(string(ch), &s)
			if err != nil {
				return -1, err
			}
			s.push(val)
		}
	}

	return s.pop().(int), nil
}
