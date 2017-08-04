package validate_test

import (
	"log"

	"github.com/northbright/validate"
)

func ExampleValidMobilePhoneNum() {
	log.Printf("---------- ValidMobilePhoneNum() Test Begin -------------")

	nums := []string{
		"aaabc89232",
		"10000",
		"13800138000",
	}

	for _, v := range nums {
		valid := validate.ValidMobilePhoneNum(v)
		log.Printf("%v: %v", v, valid)
	}

	log.Printf("---------- ValidMobilePhoneNum() Test End -------------")
	// Output:
}

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
	log.Printf("Len: %v - %v, has num: %v, has '-': %v, has '_': %v", validate.UsernameMinLen, validate.UsernameMaxLen, validate.UsernameHasNum, validate.UsernameHasHyphen, validate.UsernameHasUnderscore)

	for _, v := range usernames {
		log.Printf("%v: %v", v, validate.ValidUsername(v))
	}
	log.Printf("Test 1: Done")

	// Customized Username Validation Configuration.
	validate.UsernameMinLen = 4
	validate.UsernameHasHyphen = false
	validate.UsernameHasUnderscore = false
	log.Printf("Test 2: Customized Username Validation Configuration:")
	log.Printf("Len: %v - %v, has num: %v, has '-': %v, has '_': %v", validate.UsernameMinLen, validate.UsernameMaxLen, validate.UsernameHasNum, validate.UsernameHasHyphen, validate.UsernameHasUnderscore)

	for _, v := range usernames {
		log.Printf("%v: %v", v, validate.ValidUsername(v))
	}
	log.Printf("Test 2: Done")

	log.Printf("---------- ValidUsername() Test End -------------")
	// Output:
}

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
	log.Printf("Len: %v - %v, one num: %v, one upper: %v, one lower: %v, one special: %v", validate.PasswordMinLen, validate.PasswordMaxLen, validate.PasswordAtLeastOneNum, validate.PasswordAtLeastOneUpper, validate.PasswordAtLeastOneLower, validate.PasswordAtLeastOneSpecial)

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
	validate.PasswordMinLen = 6
	validate.PasswordAtLeastOneUpper = false
	validate.PasswordAtLeastOneSpecial = false
	log.Printf("Test 2. Customized Password Validation Configuration:")
	log.Printf("Len: %v - %v, one num: %v, one upper: %v, one lower: %v, one special: %v", validate.PasswordMinLen, validate.PasswordMaxLen, validate.PasswordAtLeastOneNum, validate.PasswordAtLeastOneUpper, validate.PasswordAtLeastOneLower, validate.PasswordAtLeastOneSpecial)

	for _, v := range passwords {
		valid, err := validate.ValidPassword(v)
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
