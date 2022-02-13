package helper

import (
	"errors"
	"regexp"
	"strings"
)

func CheckEmailPattern(email string) error {
	splitEmail := strings.Split(email, "@")

	if len(splitEmail) != 2 {
		return errors.New("email must contain exactly one local and domain name")
	}

	if strings.HasPrefix(splitEmail[0], ".") || strings.HasSuffix(splitEmail[0], ".") {
		return errors.New("local name cannot start or end with dot")
	}

	if re := regexp.MustCompile(`\.\.`); len(re.Find([]byte(splitEmail[0]))) != 0 {
		return errors.New("local name cannot contain consecutive dots")
	}

	if re := regexp.MustCompile("[^a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]"); re.MatchString(splitEmail[0]) {
		return errors.New("local name cannot contain forbidden characters")
	}

	if strings.HasPrefix(splitEmail[1], "-") || strings.HasSuffix(splitEmail[1], "-") {
		return errors.New("domain name cannot start or end with hyphen")
	}

	if strings.HasPrefix(splitEmail[1], ".") || strings.HasSuffix(splitEmail[1], ".") {
		return errors.New("domain name cannot start or end with dot")
	}

	if strings.ContainsAny(splitEmail[1], "_") {
		return errors.New("domain name cannot contain underscore")
	}

	if re := regexp.MustCompile("[^a-zA-Z0-9.-]"); re.MatchString(splitEmail[1]) {
		return errors.New("domain name cannot contain forbidden characters")
	}

	splitDomain := strings.Split(splitEmail[1], ".")

	if len(splitDomain) < 2 {
		return errors.New("domain name must contain top domain")
	}

	return nil
}
