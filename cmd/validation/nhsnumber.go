package validation

import (
	"errors"
	"strconv"

	"github.com/Joshswooft/nhs/cmd/validation/utils"
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

	validCheckSum := utils.Checksum(id)

	if !validCheckSum {
		return errors.New("nhs number is invalid")
	}

	return nil
}
