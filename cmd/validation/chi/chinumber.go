/**
Community Health Index (CHI) Number uniquely identifies a patient within the NHS
in scotland.

A CHI number is a unique, ten-digit identifier assigned to each patient on the index.

*/
package chi

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Joshswooft/nhs/cmd/validation/utils"
)

var (
	ErrChiLength          = errors.New("CHI number is not 10 characters long")
	ErrChiNonDigits       = errors.New("CHI number contains non digits")
	ErrChiInvalidDate     = errors.New("CHI number date is invalid")
	ErrChiInvalidChecksum = errors.New("CHI number checksum is invalid")
)

/** Validate checks whether the given id is a valid community health index (CHI) number
The first six digits of a CHI number are a patient's date of birth in DD/MM/YY format.
The first digit of a CHI number must, therefore, be 3 or less.

The 9th digit refers to the gender of the person: even for female, odd for male.

Last digit is for checksum.
*/
func Validate(id string) error {
	if len(id) != 10 {
		return ErrChiLength
	}
	_, err := strconv.Atoi(id)

	if err != nil {
		return ErrChiNonDigits
	}

	// DD/MM/YY
	dateStr := id[0:2] + "/" + id[2:4] + "/" + id[4:6]
	fmt.Println(dateStr)

	_, err = time.Parse("02/01/06", dateStr)

	if err != nil {
		return ErrChiInvalidDate
	}

	validCheckSum := utils.Checksum(id)

	if !validCheckSum {
		return ErrChiInvalidChecksum
	}

	return nil
}

// GetGender from a chi number id
// evaluates the 9th digit and if even => Female, odd => Male
func GetGender(id string) (utils.Gender, error) {
	if len(id) != 10 {
		return utils.Female, ErrChiLength
	}
	_, err := strconv.Atoi(id)

	if err != nil {
		return utils.Female, ErrChiNonDigits
	}

	digit := rune(id[8] - '0')

	if digit%2 == 0 {
		return utils.Female, nil
	} else {
		return utils.Male, nil
	}

}

func GetDateOfBirth(id string) (time.Time, error) {
	if len(id) != 10 {
		return time.Time{}, ErrChiLength
	}
	_, err := strconv.Atoi(id)

	if err != nil {
		return time.Time{}, ErrChiNonDigits
	}

	// DD/MM/YY
	dateStr := id[0:2] + "/" + id[2:4] + "/" + id[4:6]
	fmt.Println(dateStr)

	date, err := time.Parse("02/01/06", dateStr)

	if err != nil {
		return time.Time{}, ErrChiInvalidDate
	}
	return date, nil
}
