package uniq

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckSuccess(t *testing.T) {
	s := `I love music.
He love music.
We love music.

I love music of Kartik.
U love music of Kartik.
Thanks.
I love music of Kartik.
Dont love music of Kartik.
`
	ss := strings.SplitAfter(s, "\n")
	a := Args{
		Count:           false,
		Duplicates:      false,
		Uniq:            false,
		CaseInsensitive: false,
		FieldsNum:       0,
		CharsNum:        0,
	}

	res := Uniq(ss, a)
	fmt.Println(res)
}

func TestCheckFail(t *testing.T) {

}
