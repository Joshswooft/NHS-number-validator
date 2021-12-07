package validation

import (
	"errors"
	"strconv"
)

// NhsNumberValidator checks whether the id is a valid NHS number.
//
// NHS number is a 10 digit code that is used as a primary identifier for a patient
// within the NHS in England and Wales.
//
// NHS number is a 10 digit code that conforms to the following algorithm:
//
// 1. Multiply the 1st 9 digits by a weighting factor: 11 - index. (where index starts from 1 and goes to 9).
//
// 2. Add the results of each multiplication together
//
// 3. Divide by 11 to get the remainder
//
// 4. Subtract the remainder from 11 to get the check digit. If the result is 11 then 0 is used as the check digit.
//
// 5. Check the remainder matches the check digit. If it does not then the NHS number is invalid.
//
// The 10th digit is used as a check digit to confirm the validity of the id.
func NhsNumberValidator(id string) error {
	if len(id) != 10 {
		return errors.New("nhs number must be 10 digits")
	}

	_, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	// Step1: multiply each of the first 9 numbers by 11-index.
	factors, checkDigit := multiplyByWeightingFactor(id)

	// Step2: Add together
	total := 0
	for _, f := range factors {
		total += f
	}

	// Step3: Divide by 11 and check remainder against check digit

	remainder := total % 11

	expectedCheckDigit := 11 - remainder

	if expectedCheckDigit == 11 {
		expectedCheckDigit = 0
	}

	if checkDigit != expectedCheckDigit {
		return errors.New("nhs number is invalid")
	}

	return nil
}

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
