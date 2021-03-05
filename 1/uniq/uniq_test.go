package uniq

import (
	"reflect"
	"strings"
	"testing"
)

var args = Args{
	Count:           false,
	Duplicates:      false,
	Uniq:            false,
	CaseInsensitive: false,
	FieldsNum:       0,
	CharsNum:        0,
}

func resetArgs() {
	args = Args{
		Count:           false,
		Duplicates:      false,
		Uniq:            false,
		CaseInsensitive: false,
		FieldsNum:       0,
		CharsNum:        0,
	}
}

func TestNoParams(t *testing.T) {
	s := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

	ss2 := `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.`

	ss := strings.SplitAfter(s, "\n")
	res := Uniq(ss, args)

	expect := strings.SplitAfter(ss2, "\n")
	if !reflect.DeepEqual(res, expect) {
		t.Fatal("Unexpected result")
	}
}

func TestCflag(t *testing.T) {
	defer resetArgs()

	args.Count = true

	s := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

	s2 := `3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.`

	ss := strings.SplitAfter(s, "\n")
	res := Uniq(ss, args)

	expect := strings.SplitAfter(s2, "\n")
	if !reflect.DeepEqual(res, expect) {
		t.Fatal("Unexpected result")
	}
}

func TestDflag(t *testing.T) {
	defer resetArgs()

	args.Duplicates = true

	s := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

	s2 := `I love music.
I love music of Kartik.
I love music of Kartik.`

	ss := strings.SplitAfter(s, "\n")
	res := Uniq(ss, args)

	expect := strings.SplitAfter(s2, "\n")
	if !reflect.DeepEqual(res, expect) {
		t.Fatal("Unexpected result")
	}
}

func TestUflag(t *testing.T) {
	defer resetArgs()

	args.Uniq = true

	s := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

	s2 := `
Thanks.`

	ss := strings.SplitAfter(s, "\n")
	res := Uniq(ss, args)

	expect := strings.SplitAfter(s2, "\n")
	if !reflect.DeepEqual(res, expect) {
		t.Fatal("Unexpected result")
	}
}

func TestIflag(t *testing.T) {
	defer resetArgs()

	args.CaseInsensitive = true

	s := `I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.`

	s2 := `I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.`

	ss := strings.SplitAfter(s, "\n")
	res := Uniq(ss, args)

	expect := strings.SplitAfter(s2, "\n")
	if !reflect.DeepEqual(res, expect) {
		t.Fatal("Unexpected result")
	}
}

func TestMultiParams(t *testing.T) {
	defer resetArgs()

	args.Count = true
	args.CaseInsensitive = true
	args.FieldsNum = 2
	args.CharsNum = 3

	s := `a b xxxI LOVE MUSIC.
a b xxxI love music.
s v zzzI LoVe MuSiC.

a b xxxI love MuSIC of Kartik.
a b xxxI love music of kartik.
a b xxxThanks.
a b xxxI love music of kartik.
a b xxxI love MuSIC of Kartik.`

	s2 := `3 a b xxxI LOVE MUSIC.
1 
2 a b xxxI love MuSIC of Kartik.
1 a b xxxThanks.
2 a b xxxI love music of kartik.`

	ss := strings.SplitAfter(s, "\n")
	res := Uniq(ss, args)

	expect := strings.SplitAfter(s2, "\n")
	if !reflect.DeepEqual(res, expect) {
		t.Fatal("Unexpected result")
	}
}

