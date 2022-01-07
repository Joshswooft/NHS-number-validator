package utils

// multiplyByWeightingFactor - multiplies each of the first 9 numbers by 11 - index. (where index starts from 1)
// returns a 9 length array of ints and the check digit used for validity.
func multiplyByWeightingFactor(id string) (factors []int, checkDigit int) {
	digits := id[0:9]
	d := rune(id[9] - '0')
	checkDigit = int(d)

	factors = make([]int, len(digits))
	for i, r := range digits {
		// index starts from 0 in go
		weightingFactor := 11 - (i + 1)
		digit := int(r - '0')
		factors[i] = digit * weightingFactor
	}
	return
}

// Checksum uses the weighting factors and confirms whether the computed checksum equals the checkDigit
func Checksum(id string) bool {
	factors, checkDigit := multiplyByWeightingFactor(id)

	total := 0
	for _, f := range factors {
		total += f
	}

	remainder := total % 11

	expectedCheckDigit := 11 - remainder

	if expectedCheckDigit == 11 {
		expectedCheckDigit = 0
	}

	return checkDigit == expectedCheckDigit
}
