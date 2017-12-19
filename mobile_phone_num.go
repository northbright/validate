package validate

import (
	"regexp"
)

// ValidMobilePhoneNum validates the mobile phone number(China).
//
// Return:
//     true for valid or false for invalid.
func ValidMobilePhoneNum(number string) bool {
	p := `^\d{11}$`
	r := regexp.MustCompile(p)

	if r.MatchString(number) {
		return true
	}
	return false
}
