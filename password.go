package validate

import (
	"fmt"
	"regexp"
)

// PasswordOption specifies an option for password validation.
type PasswordOption struct {
	f func(*passwordOptions)
}

// passwordOptions represents the internal options for password validation.
type passwordOptions struct {
	minLen     int
	maxLen     int
	oneNum     bool
	oneUpper   bool
	oneLower   bool
	oneSpecial bool
}

// PasswordMinLen specifies the min length of password.
func PasswordMinLen(l int) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.minLen = l
	}}
}

// PasswordMaxLen specifies the max length of password.
func PasswordMaxLen(l int) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.maxLen = l
	}}
}

// PasswordOneNum specifies if password must have at least one number.
func PasswordOneNum(flag bool) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.oneNum = flag
	}}
}

// PasswordOneUpper specifies if password must have at least one upper case letter.
func PasswordOneUpper(flag bool) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.oneUpper = flag
	}}
}

// PasswordOneLower specifies if password must have at least one lower case letter.
func PasswordOneLower(flag bool) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.oneLower = flag
	}}
}

// PasswordOneSpecial specifies if password must have at least one special letter.
// One special letter may be one symbol or one punctuation.
func PasswordOneSpecial(flag bool) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.oneSpecial = flag
	}}
}

// ValidPassword validates the password.
//
// Params:
//     password: password string to validate.
//               Password consists of numbers, letters(lower or upper) and all chars other than numbers, letters and "_".
//     options: users can specify options by following functions:
//              PasswordMinLen(): min len. Default: 8.
//              PasswordMaxLen(): max len. Default: 64.
//              PasswordOneNum(): at least one number. Default: false.
//              PasswordOneUpper(): at least one upper case letter. Default: false.
//              PasswordOneLower(): at least one lower case letter. Default: false.
//              PasswordOneSpecial(): at least one symbol or punctuation. Default: false.
func ValidPassword(password string, options ...PasswordOption) bool {
	var patterns []string

	// Initialize default password options.
	op := passwordOptions{
		minLen: 8,
		maxLen: 64,
	}

	// Override default options with user customized options.
	for _, option := range options {
		option.f(&op)
	}

	// Password consists of numbers, letters(lower or upper) and all chars other than numbers, letters and "_".
	p := fmt.Sprintf(`^(\d|[a-z]|[A-Z]|\W){%v,%v}$`, op.minLen, op.maxLen)
	patterns = append(patterns, p)

	// Because that Golang's regexp package does NOT support `(?=re)`:
	// Zero-width positive lookahead assertion,
	// Use 4 patterns to valid the password.
	if op.oneNum {
		patterns = append(patterns, `^.*\d`)
	}

	if op.oneLower {
		patterns = append(patterns, `^.*[a-z]`)
	}

	if op.oneUpper {
		patterns = append(patterns, `^.*[A-Z]`)
	}

	if op.oneSpecial {
		patterns = append(patterns, `^.*\W`)
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if !re.MatchString(password) {
			return false
		}
	}

	return true
}
