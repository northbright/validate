package validate

import (
	"regexp"
)

func ValidIDCardNo(IDCardNo string) bool {
	patterns := []string{
		// 15-digit China ID Card Number
		`^[1-9]\d{5}\d{2}(0[1-9]|10|11|12)([0-2][1-9]|10|20|30|31)\d{3}$`,
		// 18-digit China ID Card Number
		`^[1-9]\d{5}(18|19|[23]\d)\d{2}(0[1-9]|10|11|12)([0-2][1-9]|10|20|30|31)\d{3}[0-9xX]$`,
	}

	for _, p := range patterns {
		r := regexp.MustCompile(p)

		if r.MatchString(IDCardNo) {
			return true
		}
	}
	return false
}
