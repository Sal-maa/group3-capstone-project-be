package helper

import (
	"errors"
	"strings"
)

func CheckStringInput(s string) error {
	if strings.Contains(strings.ReplaceAll(s, " ", ""), ";--") {
		return errors.New("input cannot contain forbidden character")
	}

	return nil
}
