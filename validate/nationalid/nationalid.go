package nationalid

import (
	"errors"
	"strconv"
	"strings"
)

//IsValid Check national id validation
func IsValid(value string) (bool, error) {
	valueCount := len(value)
	if valueCount < 8 {
		return false, errors.New("nationalid length must be 10")
	}

	// check len and add zero
	if valueCount >= 8 && valueCount < 10 {
		// add zero 0...
		value = strings.Repeat("0", 10-valueCount) + value
	}

	// convert string to int slice
	var valueSlices []uint8
	for _, s := range value {
		i, err := strconv.ParseInt(string(s), 10, 8)
		if err != nil {
			return false, errors.New("value problem")
		}
		valueSlices = append(valueSlices, uint8(i))
	}

	// check #1 > *position
	s, err := calculateNationalIDNumbers(&valueSlices)
	if err != nil {
		return true, err
	}

	// check #2
	s %= 11
	l := valueSlices[len(value)-1]
	if (s < 2 && s != int(l)) || (s >= 2 && int(l) != 11-s) {
		return false, errors.New("nationalid is not valid")
	}

	return true, nil
}

//calculateNationalIDNumbers Yo
func calculateNationalIDNumbers(valueSlices *[]uint8) (sum int, err error) {
	var i uint8
	for i = 10; i >= 2; i-- {
		sum += int(i * (*valueSlices)[10-i])
	}
	return sum, nil
}
