package calc

import (
	"fmt"
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

func TestCalc(t *testing.T) {
	for _, tt := range expressiontests {
		res := Calc(InfToPosf(tt.in))
		fmt.Println(tt.in, " >>> ", res)
		require.Equal(t, tt.out, res)
	}
}
