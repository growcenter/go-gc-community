package validate

import (
	"regexp"
)

func Email(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}

func PhoneNumber(phone string) bool {
	if len(phone) < 10 || len(phone) > 13 {
		return false
	}
	
	pattern := `^(0|\\+62|062|62)[0-9]+$`

	return regexp.MustCompile(pattern).MatchString(phone)
}