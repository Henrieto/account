package validators

import (
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	// This regex might not cover all edge cases but works for most standard email addresses
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

const (
	UPPPERCASE = "ABCDEFGHIJKLMNOPQSRTUVWXYZ"
	LOWERCASE  = "abcdefghijklmnopqrstuvwxyz"
	CHARACTERS = "`~!@#$%^&*()-_=+[{]}/>.<,]"
)

// method to check if the password is strong enough
// and also if the password contains the required characters
func IsStrongPassword(password string) error {
	if len(password) < 8 {
		return ErrorPasswordNotStrong
	}
	contains_uppercase := strings.ContainsAny(UPPPERCASE, password)
	contains_lowercase := strings.ContainsAny(LOWERCASE, password)
	contains_characters := strings.ContainsAny(CHARACTERS, password)

	if !contains_uppercase {
		return ErrorPasswordNotStrong
	}
	if !contains_lowercase {
		return ErrorPasswordNotStrong
	}
	if !contains_characters {
		return ErrorPasswordNotStrong
	}
	return nil
}
