package calc

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

var expressiontests = []struct {
	in  string
	out int
}{
	{"3*(2+1)/((1+2)*3)+6", 7},
	{"23+(((((((((9)*8)*7)*6)*5)*4)*3)*2)*1)", 362903},
	{"(1+2)-3", 0},
	{"(1+2)*3", 9},
	{"(12 + 25) - 36", 1},
	{"10+15*22", 340},
	{"25-(64*3-(57-123)*(31+4))*(27-11)", -40007},
	{"12224/2-(3*6)+(2-(123/123)) ", 6095},
	{"10-9-8-7-6-5-4-3-2-1", -35},
	{"10*5*2*1", 100},
}

var errortests = []struct {
	in string
	err error
}{
	{"1)+783*(23) -(1", errors.New("bad syntax: parenthesis don't match")},
	{"(10-1)*(8 * 356 * (174 - (15 + 1))",  errors.New("bad syntax: parenthesis don't match")},
}

var operrortests = []struct {
	in string
	err error
}{
	{"3^2", errors.New("bad syntax: unsupported operation")},
	{"123 @ (3956 + 534) & 1",  errors.New("bad syntax: unsupported operation")},
}


func TestCalc(t *testing.T) {
	for _, tt := range expressiontests {
		exp, _ := InfToPosf(tt.in)
		res, _ := Calc(exp)
		require.Equal(t, tt.out, res)
	}
}

func TestErrors(t *testing.T) {
	for _, tt := range errortests {
		_, err := InfToPosf(tt.in)
		require.Equal(t, tt.err, err)
	}
}

func TestOpError(t *testing.T) {
	for _, tt := range operrortests {
		exp, _ := InfToPosf(tt.in)
		_, err := Calc(exp)
		require.Equal(t, tt.err, err)
	}
}
