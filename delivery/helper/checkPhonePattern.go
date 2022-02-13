package helper

import (
	"errors"
	"regexp"
)

func CheckPhonePattern(phone string) error {
	re := regexp.MustCompile("[^0-9]")

	if re.MatchString(phone) {
		return errors.New("phone number must contain numbers only")
	}

	if len(phone) < 10 {
		return errors.New("phone number too short")
	}

	if len(phone) > 12 {
		return errors.New("phone number too long")
	}

	return nil
}
