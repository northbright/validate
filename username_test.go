package validate_test

import (
	"log"

	"github.com/northbright/validate"
)

func ExampleValidUsername() {
	log.Printf("---------- ValidUsername() Test Begin -------------")

	usernames := []string{
		"aaaa",
		"世界",
		"13800138000",
		"a__zzzz",
		"mio--cat",
		"Beyond喜欢你",
		"色褪せぬ蒼青の欠片",
		"#small!!!!!",
		"Michael.Learns.To.Rock",
		" Space Space ",
		"admin@mydomain.com",
	}

	// Default Username Validation Configuration.
	log.Printf("Test 1: Default Username Validation Configuration:")
	log.Printf("username len: 6 - 64, no dot: false, no hyphen: false, no underscore: false.")

	for _, v := range usernames {
		log.Printf("%v(len: %v): %v", v, len(v), validate.ValidUsername(v))
	}
	log.Printf("Test 1: Done")

	// Customized Username Validation Configuration.
	log.Printf("Test 2: Customized Username Validation Configuration:")
	log.Printf("username len: 8 - 64, no dot: true, no hyphen: true, no underscore: true.")

	for _, v := range usernames {
		log.Printf("%v(len: %v): %v",
			v,
			len(v),
			validate.ValidUsername(v,
				validate.UsernameMinLen(8),
				validate.UsernameNoDot(true),
				validate.UsernameNoHyphen(true),
				validate.UsernameNoUnderscore(true),
			),
		)
	}
	log.Printf("Test 2: Done")
	log.Printf("---------- ValidUsername() Test End -------------")
	// Output:
}
