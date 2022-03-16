package phone

import (
	"errors"
	"regexp"
)

// IsValid mobile phone number validation
func IsValid(value string) (bool, error) {
	re := regexp.MustCompile(`^(?:0|\+98|0098)(\d{10})$`)
	result := re.FindStringSubmatch(value)
	if len(result) > 0 {
		return true, nil
	}
	return false, errors.New("invalid phone number")
}
