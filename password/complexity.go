package password

// Meet returns true if required complexity is met
// It accepts the same rules than Generate, but it allows to set:
// minLenght ensures password is not smaller.
// maxLenght ensures password is not longer that maximum lenght.
// exact ensures password has the required amount of characters,
// if false it will check if the password has the type if num is
// greater than 0
func (g *Generator) Meet(source string, minLenght, maxLenght, numDigits, numSymbols int, noUpper, allowRepeat, exact bool) bool {

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

	if exact {
		if numDigits != digits {
			return false
		}
		if numSymbols != symbols {
			return false
		}
	} else {
		// it must have at least 1 digit
		if numDigits > 0 && digits == 0 {
			return false
		}

		// it must have at least 1 symbol
		if numSymbols > 0 && symbols == 0 {
			return false
		}
	}

	// upper not requested but there are
	if noUpper && upper > 0 {
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
