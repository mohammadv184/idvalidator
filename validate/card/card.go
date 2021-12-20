package card

import (
	"bytes"
	"errors"
	"strconv"
)

var firstNumValid = []byte{
	54, //6
	53, //5
	42, //4
}

//IsValid check card number validation
func IsValid(value string) (bool, error) {
	// check length
	if len(value) != 16 {
		return false, errors.New("card length must be 16")
	}

	// check first number
	if !bytes.Contains(firstNumValid, []byte{value[0]}) {
		return false, errors.New("first number isn`t valid")
	}

	// check calculate *slices < 10
	sum := calculateCardNumbers(&value)

	// check #3 > %10 = 0
	if sum%10 != 0 {
		return false, errors.New("card number is not valid")
	}

	return true, nil
}

//calculateCardNumbers Yo
func calculateCardNumbers(valueSlices *string) (sum uint64) {
	for i := 0; i < len(*valueSlices); i++ {
		v, _ := strconv.ParseUint(string((*valueSlices)[i]), 10, 8)
		var x uint64
		if i%2 == 0 {
			if x = v * 2; x >= 10 {
				x -= 9
			}
		} else {
			x = v
		}
		sum += x
	}

	return sum
}
