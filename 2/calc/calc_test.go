package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalc(t *testing.T) {
	exp := "3*(2+1)/((1+2)*3)+6"
	exp = InfToPosf(exp)
	require.Equal(t, 7, Calc(exp))
}
