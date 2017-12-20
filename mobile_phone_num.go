package validate

import (
	"regexp"
)

// ValidMobilePhoneNumInChina validates the mobile phone number in China.
//
// Return:
//     true for valid or false for invalid.
func ValidMobilePhoneNumInChina(number string) bool {
	p := `^\d{11}$`
	r := regexp.MustCompile(p)

	if r.MatchString(number) {
		return true
	}
	return false
}
