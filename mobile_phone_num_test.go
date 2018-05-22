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
