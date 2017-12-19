package validate

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// UsernameOption specifies an option for username validation.
type UsernameOption struct {
	f func(*usernameOptions)
}

// usernameOptions represents the internal options for username validation.
type usernameOptions struct {
	minLen       int
	maxLen       int
	noHyphen     bool
	noUnderscore bool
}

// UsernameMinLen specifies the min length of username.
func UsernameMinLen(l int) UsernameOption {
	return UsernameOption{func(op *usernameOptions) {
		op.minLen = l
	}}
}

// UsernameMaxLen specifies the max length of username.
func UsernameMaxLen(l int) UsernameOption {
	return UsernameOption{func(op *usernameOptions) {
		op.maxLen = l
	}}
}

// UsernameNoHyphen specifies if username can have hyphens.
func UsernameNoHyphen(flag bool) UsernameOption {
	return UsernameOption{func(op *usernameOptions) {
		op.noHyphen = flag
	}}
}

// UsernameNoUnderscore specifies if username can have underscores.
func UsernameNoUnderscore(flag bool) UsernameOption {
	return UsernameOption{func(op *usernameOptions) {
		op.noUnderscore = flag
	}}
}

// ValidUsername validates the username.
//
// Params:
//     username: username to validate.
//     options: users can specifiy options by following functions:
//              UsernameMinLen(): min len. Default: 6.
//              UsernameMaxLen(): max len. Default: 64.
//              UsernameNoNum(): no number in username. Default: false.
//              UsernameNoHyphen(): no hyphen in username. Default: false.
//              UsernameNoUnderscore(): no underscore in username. Default: false.
// Return:
//     true for valid or false for invalid.
// Comments:
//     Username can contain Latin letters, digits and Chinese characters.
//     Each Chinese character's length is recognized as 2.
//     Because the font width of Chinese character is 2x than Latin(number) in most case.
//     e.g.
//     "中文汉字" -> 4 Chinese Characters: display width = 8
//     "abcd1234" -> 8 Latin chars and numbers mixed: display width = 8.5
func ValidUsername(username string, options ...UsernameOption) bool {
	// Intialize default options.
	op := usernameOptions{
		minLen: 6,
		maxLen: 64,
	}

	// Override options with user customized options.
	for _, option := range options {
		option.f(&op)
	}

	p := fmt.Sprintf(`^(?:\p{Han}|[a-zA-Z]|\d)+$`)

	if !op.noHyphen {
		p = strings.Replace(p, `)`, `|-)`, 1)
	}

	if !op.noUnderscore {
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

	if l < op.minLen || l > op.maxLen {
		return false
	}

	return true
}
