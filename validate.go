package validate

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	// Username Validation Configuration.

	// UsernameMinLen represents the min length of username.
	// Username can contain Latin letters and Chinese characters.
	// Each Chinese character's length is recognized as 2.
	// Because the font width of Chinese character is 2x than Latin(number) in most case.
	// e.g.
	// "中文汉字" -> 4 Chinese Characters: display width = 8
	// "abcd1234" -> 8 Latin chars and numbers mixed: display width = 8.5
	UsernameMinLen = 6
	// UsernameMaxLen represents the max length of username.
	// See UsernameMinLen for more information.
	UsernameMaxLen = 16
	// UsernameHasNum represents if username can have number.
	UsernameHasNum = true
	// UsernameHasHyphen represents if username can have hyphen('-').
	UsernameHasHyphen = true
	// UsernameHasUnderscore represents if username can have hyphen('_').
	UsernameHasUnderscore = true

	// Password Validation Configuration.

	// PasswordMinLen represents the min length of password.
	// The length is rune count of UTF-8 string. Ex: rune count of "Hello, 世界" is 9.
	PasswordMinLen = 8
	// PasswordMaxLen represents the max length of password.
	PasswordMaxLen = 64
	// PasswordAtLeastOneNum represents if password should have at least one number.
	PasswordAtLeastOneNum = true
	// PasswordAtLeastOneUpper represents if password should have at least one upper-case letter.
	PasswordAtLeastOneUpper = true
	// PasswordAtLeastOneLower represents if password should have at least one lower-case letter.
	PasswordAtLeastOneLower = true
	// PasswordAtLeastOneSpecial represents if password should have at least one special character.
	// One special characer may be one symbol or one punctuation.
	PasswordAtLeastOneSpecial = true

	// Errors of Invalid Password

	// ErrPasswordInvalidUTF8Rune is the reason of invalid UTF-8 rune.
	ErrPasswordInvalidUTF8Rune = fmt.Errorf("password consists of invalid UTF-8 rune")
	// ErrPasswordLen is the reason of invalid length.
	ErrPasswordLen = fmt.Errorf("invalid password length")
	// ErrPasswordAtLeastOneNum is the reason of not having at least one number.
	ErrPasswordAtLeastOneNum = fmt.Errorf("password should have at least one number")
	// ErrPasswordAtLeastOneUpper is the reason of not having at least one upper-case letter.
	ErrPasswordAtLeastOneUpper = fmt.Errorf("password should have at least one uppler-case letter")
	// ErrPasswordAtLeastOneLower is the reason of not having at least one lower-case letter.
	ErrPasswordAtLeastOneLower = fmt.Errorf("password should have at least one lower-case letter")
	// ErrPasswordAtLeastOneSpecial is the reason of not having at least one special character.
	// One special letter may be one symbol or one punctuation.
	ErrPasswordAtLeastOneSpecial = fmt.Errorf("password should have at least one special character")
)

// ValidMobilePhoneNum validates the mobile phone number.
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

// ValidUsername validates the username.
//
// Return:
//    true for valid or false for invalid.
// Comments:
//     Username can contain Latin letters and Chinese characters.
//         Each Chinese character's length is recognized as 2.
//         Because the font width of Chinese character is 2x than Latin(number) in most case.
//         e.g.
//         "中文汉字" -> 4 Chinese Characters: display width = 8
//         "abcd1234" -> 8 Latin chars and numbers mixed: display width = 8.5
//    Username validation can be configured by following variables:
//    UsernameMinLen(default: 6)
//    UsernameMaxLen(default: 64)
//    UsernameHasNum(default: true)
//    UsernameHasHyphen(default: true)
//    UsernameHasUnderscore(default: true)
func ValidUsername(username string) bool {
	p := fmt.Sprintf(`^(?:\p{Han}|[a-zA-Z])+$`)

	if UsernameHasNum {
		p = strings.Replace(p, `)`, `|\d)`, 1)
	}

	if UsernameHasHyphen {
		p = strings.Replace(p, `)`, `|-)`, 1)
	}

	if UsernameHasUnderscore {
		p = strings.Replace(p, `)`, `|_)`, 1)
	}

	re := regexp.MustCompile(p)
	if !re.MatchString(username) {
		return false
	}

	l := 0
	for _, r := range username {
		switch {
		case unicode.Is(unicode.Scripts["Han"], r):
			l += 2
		default:
			l++
		}
	}

	if l < UsernameMinLen || l > UsernameMaxLen {
		return false
	}

	return true
}

// ValidPassword validates the password.
//
// Return:
//     valid: true, nil
//     invalid: false, error with message.
//     The error may be one of the following:
//         nil: valid
//         ErrPasswordInvalidUTF8Rune
//         ErrPasswordLen
//         ErrPasswordAtLeastOneNum
//         ErrPasswordAtLeastOneUpper
//         ErrPasswordAtLeastOneLower
//         ErrPasswordAtLeastOneSpecial
// Comments:
//     Password validation can be configured by following variables:
//     PasswordMinLen(default: 8)
//     PasswordMaxLen(default: 64)
//     PasswordAtLeastOneNum(default: true)
//     PasswordAtLeastOneUpper(default: true)
//     PasswordAtLeastOneLower(default: true)
//     PasswordAtLeastOneSpecial(default: true)
func ValidPassword(password string) (bool, error) {
	var (
		oneNum     = false
		oneUpper   = false
		oneLower   = false
		oneSpecial = false
	)

	// Whether password consists entirely of valid UTF-8-encoded runes.
	if !utf8.ValidString(password) {
		return false, ErrPasswordInvalidUTF8Rune
	}

	// Validate Password Length.
	len := utf8.RuneCountInString(password)
	if len < PasswordMinLen || len > PasswordMaxLen {
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

	if PasswordAtLeastOneNum && !oneNum {
		return false, ErrPasswordAtLeastOneNum
	}

	if PasswordAtLeastOneUpper && !oneUpper {
		return false, ErrPasswordAtLeastOneUpper
	}

	if PasswordAtLeastOneLower && !oneLower {
		return false, ErrPasswordAtLeastOneLower
	}

	if PasswordAtLeastOneSpecial && !oneSpecial {
		return false, ErrPasswordAtLeastOneSpecial
	}

	return true, nil
}
