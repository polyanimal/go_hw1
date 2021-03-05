package calc

import "testing"

func TestCalc(t *testing.T) {
	exp := "3*(2+1)/((1+2)*3)+6"
	exp = InfToPosf(exp)
	if Calc(exp) != 7 {
		t.Fatal("Wrong Answer")
	}
}