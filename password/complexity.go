package password

// Meet returns true if required complexity is met
// It accepts the same rules than Generate, but it allows to set:
// minLenght ensures password is not smaller.
// maxLenght ensures password is not longer that maximum lenght.
// minLower ensures password have lowercase characters.
// minUpper ensures password have uppercase characters.
// exact ensures password has the required amount of characters,
// if true numType need to match
// if false it will check the type is used, but not the amount
func (g *Generator) Meet(source string, minLenght, maxLenght int, hasLower, hasUpper, hasDigits, hasSymbols, allowRepeat bool) bool {

	var (
		sourceBytes []byte
		digits      int
		symbols     int
		lower       int
		upper       int
		repeat      int
	)

	if len(source) < minLenght {
		return false
	}

	if (len(source)) > maxLenght {
		return false
	}

	// calculate amounts
	for i := 0; i < len(source); i++ {

		if byteInString(source[i], g.digits) {
			digits = digits + 1
		}

		if byteInString(source[i], g.symbols) {
			symbols = symbols + 1
		}

		if byteInString(source[i], g.lowerLetters) {
			lower = lower + 1
		}

		if byteInString(source[i], g.upperLetters) {
			upper = upper + 1
		}

		if byteInArray(source[i], sourceBytes) {
			repeat = repeat + 1
		} else {
			sourceBytes = append(sourceBytes, source[i])
		}

	}

	// it must have at least 1 digit
	if hasDigits && digits == 0 {
		return false
	}

	// it must have at least 1 symbol
	if hasSymbols && symbols == 0 {
		return false
	}

	// it must have at least 1 lowercase character
	if hasLower && lower == 0 {
		return false
	}

	// it must have at least 1 uppercase character
	if hasUpper && upper == 0 {
		return false
	}

	// fail if repetitions are not allowed, but there are
	// if repetitions are allowed but password does not have must be considered as valid
	if !allowRepeat && repeat > 0 {
		return false
	}
	return true

}

func byteInArray(item byte, items []byte) bool {

	for _, thisItem := range items {
		if item == thisItem {
			return true
		}
	}

	return false

}

func byteInString(item byte, source string) bool {

	for i := 0; i < len(source); i++ {
		if source[i] == item {
			return true
		}
	}

	return false
}
