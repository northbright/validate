package validate

import (
	"fmt"
	"regexp"
	"strings"
)

// UsernameOption specifies an option for username validation.
type UsernameOption struct {
	f func(*usernameOptions)
}

// usernameOptions represents the internal options for username validation.
type usernameOptions struct {
	minLen       int
	maxLen       int
	noDot        bool
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

// UsernameNoDot specifies if username can have dots.
func UsernameNoDot(flag bool) UsernameOption {
	return UsernameOption{func(op *usernameOptions) {
		op.noDot = flag
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
// Username can contain unicode letters, latin letters and digits.
//
// Params:
//     username: username to validate.
//     options: users can specifiy options by following functions:
//              UsernameMinLen(): min length of username. Default: 6.
//                  The length of username is the number of bytes in the string(UTF-8 encoded).
//                  e.g. len("世界") = 6, len("world") = 5.
//              UsernameMaxLen(): max length of username. Default: 64.
//              UsernameNoNum(): no number in username. Default: false.
//              UsernameNoHyphen(): no hyphen in username. Default: false.
//              UsernameNoUnderscore(): no underscore in username. Default: false.
// Return:
//     true for valid or false for invalid.
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

	p := fmt.Sprintf(`^(?:\p{L}|[a-zA-Z]|\d)+$`)

	if !op.noDot {
		p = strings.Replace(p, `)`, `|\.)`, 1)
	}

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

	l := len(username)
	if l < op.minLen || l > op.maxLen {
		return false
	}

	return true
}
