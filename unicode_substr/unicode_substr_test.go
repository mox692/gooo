package unicode_substr

import "testing"

type testUnicodeSubstr struct {
	input1 string
	input2 int
	expect string
}

func TestUnicodeSubstr(t *testing.T) {
	testData := []testUnicodeSubstr{
		{
			input1: "もとゆき",
			input2: 2,
			expect: "もと",
		},
		{
			input1: "fadsfa",
			input2: 100,
			expect: "fadsfa",
		},
		{
			input1: "😍😍😍",
			input2: 2,
			expect: "😍😍",
		},
		{
			input1: "﷽ﷺ	ﷺ	",
			input2: 5,
			expect: "﷽ﷺ	ﷺ	",
		},
	}

	for _, v := range testData {
		if ret := UnicodeSubstr(v.input1, v.input2); ret != v.expect {
			t.Errorf("expect %s, got %s\n", v.expect, ret)
		}
	}
}
