package validate_test

import (
	"log"

	"github.com/northbright/validate"
)

func ExampleValidPassword() {
	log.Printf("---------- ValidPassword() Test Begin -------------")

	passwords := []string{
		"aaa123",
		"Password12",
		"Password2@",
		"aaaabbbb",
		"#ABCD1234",
		"@5431efgh",
		"Copy&Paste中文密码123",
		"Copy&Paste日本のパスワード123",
	}

	// Default Password Validation Configuration.
	log.Printf("Test 1: Default Password Validation Configuration:")
	log.Printf("len: 8 - 64, one num: false, one upper: false, one lower: false, one special: false")

	for _, v := range passwords {
		valid := validate.ValidPassword(v)
		log.Printf("%v(len: %v): %v", v, len(v), valid)
	}
	log.Printf("Test 1: Done")

	// Customized Password Validation Configuration.
	log.Printf("Test 2. Customized Password Validation Configuration:")
	log.Printf("len: 9 - 64, one num: true, one upper: true, one lower: true, one special: true")

	for _, v := range passwords {
		valid := validate.ValidPassword(v,
			validate.PasswordMinLen(9),
			validate.PasswordOneNum(true),
			validate.PasswordOneUpper(true),
			validate.PasswordOneLower(true),
			validate.PasswordOneSpecial(true),
		)
		log.Printf("%v(len: %v): %v", v, len(v), valid)
	}
	log.Printf("Test 2: Done")
	log.Printf("---------- ValidPassword() Test End -------------")
	// Output:
}
