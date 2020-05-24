package password

import (
	"testing"
)

func TestComplexity(t *testing.T) {

	gen, err := NewGenerator(nil)
	if err != nil {
		t.Fatal(err)
	}

	type c struct {
		Length  int
		Digits  int
		Symbols int
		NoUpper bool
		Repeat  bool
	}

	var cases []c

	for _, d := range []int{0, 6} {
		for _, s := range []int{0, 6} {
			for _, u := range []bool{true, false} {
				for _, r := range []bool{true, false} {
					cases = append(cases, c{
						Length:  20,
						Digits:  d,
						Symbols: s,
						NoUpper: u,
						Repeat:  r,
					})
				}
			}
		}
	}

	for _, this := range cases {
		for i := 0; i < 4; i++ {

			res, err := gen.Generate(this.Length, this.Digits, this.Symbols, this.NoUpper, this.Repeat)
			if err != nil {
				t.Error(err)
			}

			for _, exact := range []bool{true, false} {
				if expected := gen.Meet(res, this.Length-2, this.Length+2, this.Digits, this.Symbols, this.NoUpper, this.Repeat, exact); !expected {
					t.Errorf("expected %t to be %t", expected, true)
				}
			}
		}
	}

	// fail because of min length
	if expected := gen.Meet("aaa", 8, 8, 0, 0, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of max length
	if expected := gen.Meet("123456789012", 8, 8, 0, 0, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numDigits
	if expected := gen.Meet("abcd", 2, 8, 1, 0, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numDigits & exact
	if expected := gen.Meet("abcd123", 2, 8, 1, 0, false, false, true); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numDigits
	if expected := gen.Meet("abcd", 2, 8, 0, 1, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numDigits & exact
	if expected := gen.Meet("abcd$%&", 2, 8, 0, 1, false, false, true); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of noUpper
	if expected := gen.Meet("abcdABC", 2, 8, 0, 0, true, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of repetitions
	if expected := gen.Meet("a1$a1$", 2, 8, 0, 0, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

}
