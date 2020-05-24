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
							Length:  20,
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
		for i := 0; i < 4; i++ {
			var (
				res string
				err error
			)

			res, err = gen.Generate(this.Length, this.Digits, this.Symbols, true, this.Repeat)

			if err != nil {
				t.Error(err)
			}

			for _, exact := range []bool{true, false} {
				if expected := gen.Meet(res, this.Length-2, this.Length+2, this.Lower, 0, this.Digits, this.Symbols, this.Repeat, exact); !expected {
					fmt.Println(expected)
					t.Errorf("expected %t to be %t", expected, true)
				}
			}
		}
	}

	// fail because of min length
	if expected := gen.Meet("aaa", 8, 8, 0, 0, 0, 0, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of max length
	if expected := gen.Meet("123456789012", 8, 8, 0, 0, 0, 0, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numDigits
	if expected := gen.Meet("abcd", 2, 8, 0, 0, 1, 0, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numDigits & exact
	if expected := gen.Meet("abcd123", 2, 8, 0, 0, 1, 0, false, true); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numSymbols
	if expected := gen.Meet("abcd", 2, 8, 0, 0, 0, 1, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numSymbols & exact
	if expected := gen.Meet("abcd$%&", 2, 8, 0, 0, 0, 1, false, true); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numLower
	if expected := gen.Meet("ABCD", 2, 8, 1, 0, 0, 0, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numLower & exact
	if expected := gen.Meet("ABCD$%&", 2, 8, 1, 0, 0, 0, false, true); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// fail because of numUpper
	if expected := gen.Meet("abcd", 2, 8, 0, 1, 0, 0, false, false); expected {
		t.Errorf("expected %t to be %t", expected, false)
	}

	// // fail because of numUpper & exact
	// if expected := gen.Meet("abcd$%&", 2, 8, 0, 1, 0, 0, false, true); expected {
	// 	t.Errorf("expected %t to be %t", expected, false)
	// }
	// // true because of numUpper & exact
	// if expected := gen.Meet("abcd$%&A", 2, 8, 0, 1, 0, 0, false, true); expected {
	// 	t.Errorf("expected %t to be %t", expected, true)
	// }

	// // fail because of repetitions
	// if expected := gen.Meet("a1$a1$", 2, 8, 0, 0, 0, 0, false, false); expected {
	// 	t.Errorf("expected %t to be %t", expected, false)
	// }

}
