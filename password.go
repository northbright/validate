package validate

import (
	"unicode"
	"unicode/utf8"
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
//     options: users can specify options by following functions:
//              PasswordMinLen(): min len. Default: 8.
//              PasswordMaxLen(): max len. Default: 64.
//              PasswordOneNum(): at least one number. Default: false.
//              PasswordOneUpper(): at least one upper case letter. Default: false.
//              PasswordOneLower(): at least one lower case letter. Default: false.
//              PasswordOneSpecial(): at least one special letter(one symbol or one punctuation).
func ValidPassword(password string, options ...PasswordOption) bool {
	var (
		oneNum     = false
		oneUpper   = false
		oneLower   = false
		oneSpecial = false
	)

	// Initialize default password options.
	op := passwordOptions{
		// minLen represents the min length of password.
		// The length is rune count of UTF-8 string. Ex: rune count of "Hello, 世界" is 9.
		minLen: 8,
		maxLen: 64,
	}

	// Override default options with user customized options.
	for _, option := range options {
		option.f(&op)
	}

	// Whether password consists entirely of valid UTF-8-encoded runes.
	if !utf8.ValidString(password) {
		return false
	}

	// Validate Password Length.
	len := utf8.RuneCountInString(password)
	if len < op.minLen || len > op.maxLen {
		return false
	}

	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			oneNum = true
		case unicode.IsUpper(r):
			oneUpper = true
		case unicode.IsLower(r):
			oneLower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			oneSpecial = true
		}
	}

	if op.oneNum && !oneNum {
		return false
	}

	if op.oneUpper && !oneUpper {
		return false
	}

	if op.oneLower && !oneLower {
		return false
	}

	if op.oneSpecial && !oneSpecial {
		return false
	}

	return true
}
