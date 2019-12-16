package lesson3

import "testing"

func TestNone(t *testing.T) {
	t.Log("Hello, World!")
}

func TestUnpackEqual(t *testing.T) {
	res, err := Unpack(`foo`);
	if err != nil {
		t.Error(err)
	}
	if "foo" != res {
		t.Error(`foo != foo`)
	}
}

type Test struct {
	given string
	expect string
}

func TestUnpack(t *testing.T) {

	cases := [7]Test {
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45",  ""},
		{"", ""},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for i := 0; i < len(cases); i +=1 {
		t.Run("case" + string(i), func(t *testing.T) {
			result, _ := Unpack(cases[i].given)
			//if err != nil {
			//	t.Error(err)
			//}
			if cases[i].expect != result {
				t.Errorf(`'%s': expect '%s' but got '%s'`, cases[i].given, cases[i].expect, result)
			}
		})
	}
}