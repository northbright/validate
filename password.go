package validate

import (
	"fmt"
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

var (
	// ErrPasswordInvalidUTF8Rune is the reason of invalid UTF-8 rune.
	ErrPasswordInvalidUTF8Rune = fmt.Errorf("password consists of invalid UTF-8 rune")
	// ErrPasswordLen is the reason of invalid length.
	ErrPasswordLen = fmt.Errorf("invalid password length")
	// ErrPasswordOneNum is the reason of not having at least one number.
	ErrPasswordOneNum = fmt.Errorf("password should have at least one number")
	// ErrPasswordOneUpper is the reason of not having at least one upper-case letter.
	ErrPasswordOneUpper = fmt.Errorf("password should have at least one uppler-case letter")
	// ErrPasswordOneLower is the reason of not having at least one lower-case letter.
	ErrPasswordOneLower = fmt.Errorf("password should have at least one lower-case letter")
	// ErrPasswordOneSpecial is the reason of not having at least one special character.
	// One special letter may be one symbol or one punctuation.
	ErrPasswordOneSpecial = fmt.Errorf("password should have at least one special character")
)

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

// PasswordOneNum specifies if password must have at least one upper case letter.
func PasswordOneUpper(flag bool) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.oneUpper = flag
	}}
}

// PasswordOneNum specifies if password must have at least one lower case letter.
func PasswordOneLower(flag bool) PasswordOption {
	return PasswordOption{func(op *passwordOptions) {
		op.oneLower = flag
	}}
}

// PasswordOneNum specifies if password must have at least one special letter.
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
//     options: users can specifiy options by following functions:
//              PasswordMinLen(): min len. Default: 8.
//              PasswordMaxLen(): max len. Default: 64.
//              PasswordOneNum(): at least one number. Default: false.
//              PasswordOneUpper(): at least one upper case letter. Default: false.
//              PasswordOneLower(): at least one lower case letter. Default: false.
//              PasswordOneSpecial(): at least one special letter(one symbol or one punctuation).
// Return:
//     valid: true, nil
//     invalid: false, error with message.
//     The error may be one of the following:
//         nil: valid
//         ErrPasswordInvalidUTF8Rune
//         ErrPasswordLen
//         ErrPasswordOneNum
//         ErrPasswordOneUpper
//         ErrPasswordOneLower
//         ErrPasswordOneSpecial
func ValidPassword(password string, options ...PasswordOption) (bool, error) {
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
		return false, ErrPasswordInvalidUTF8Rune
	}

	// Validate Password Length.
	len := utf8.RuneCountInString(password)
	if len < op.minLen || len > op.maxLen {
		return false, ErrPasswordLen
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
		return false, ErrPasswordOneNum
	}

	if op.oneUpper && !oneUpper {
		return false, ErrPasswordOneUpper
	}

	if op.oneLower && !oneLower {
		return false, ErrPasswordOneLower
	}

	if op.oneSpecial && !oneSpecial {
		return false, ErrPasswordOneSpecial
	}

	return true, nil
}
