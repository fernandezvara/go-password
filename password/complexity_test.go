package password

import (
	"fmt"
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
		Lower   int
		Upper   int
		Repeat  bool
	}

	var cases []c

	for _, d := range []int{0, 4} {
		for _, s := range []int{0, 4} {
			for _, l := range []int{0, 4} {
				for _, u := range []int{0, 4} {
					for _, r := range []bool{true, false} {
						cases = append(cases, c{
							Length:  40,
							Digits:  d,
							Symbols: s,
							Lower:   l,
							Upper:   u,
							Repeat:  r,
						})
					}
				}
			}
		}
	}

	for _, this := range cases {
		for i := 0; i < 100; i++ {
			var (
				res string
				err error
			)

			res, err = gen.Generate(this.Length, this.Digits, this.Symbols, false, this.Repeat)
			if err != nil {
				t.Error(err)
			}

			if expected := gen.Meet(res, this.Length-2, this.Length+2, this.Lower != 0, this.Upper != 0, this.Digits != 0, this.Symbols != 0, this.Repeat); !expected {
				fmt.Println(expected)
				t.Errorf("expected %t to be %t", expected, true)
			}
		}
	}

	// fail because of min length
	if expected := gen.Meet("aaa", 8, 8, false, false, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of max length
	if expected := gen.Meet("123456789012", 8, 8, false, false, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of digits
	if expected := gen.Meet("abcd", 2, 8, false, false, true, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of symbols
	if expected := gen.Meet("abcd", 2, 8, false, false, false, true, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of lower
	if expected := gen.Meet("ABCD", 2, 8, true, false, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numUpper
	if expected := gen.Meet("abcd", 2, 8, false, true, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of repetitions (no allowed)
	if expected := gen.Meet("a1$a1$", 2, 8, false, false, false, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of repetitions (allowed)
	if expected := gen.Meet("a1$a1$", 2, 8, false, false, false, false, true); !expected {
		t.Errorf("expected %t to be %t", expected, true)
	}

}
