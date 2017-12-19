package validate_test

import (
	"log"

	"github.com/northbright/validate"
)

func ExampleValidPassword() {
	log.Printf("---------- ValidPassword() Test Begin -------------")

	passwords := []string{
		"aaa123",
		"Password1",
		"aaaabbbb",
		"#ABCD1234",
		"@5431efgh",
		"Copy&Paste中文密码123",
		"Copy&Paste日本のパスワード123",
	}

	// Default Password Validation Configuration.
	log.Printf("Test 1: Default Password Validation Configuration:")

	for _, v := range passwords {
		valid, err := validate.ValidPassword(v)
		if err != nil {
			log.Printf("%v: %v", v, err)
		} else {
			log.Printf("%v: %v", v, valid)
		}
	}
	log.Printf("Test 1: Done")

	// Customized Password Validation Configuration.
	log.Printf("Test 2. Customized Password Validation Configuration:")

	for _, v := range passwords {
		valid, err := validate.ValidPassword(v,
			validate.PasswordMinLen(6),
			validate.PasswordOneUpper(false),
			validate.PasswordOneSpecial(false),
		)
		if err != nil {
			log.Printf("%v: %v", v, err)
		} else {
			log.Printf("%v: %v", v, valid)
		}
	}
	log.Printf("Test 2: Done")
	log.Printf("---------- ValidPassword() Test End -------------")
	// Output:
}
