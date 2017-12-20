package validate_test

import (
	"log"

	"github.com/northbright/validate"
)

func ExampleValidMobilePhoneNumInChina() {
	log.Printf("---------- ValidMobilePhoneNumInChina() Test Begin -------------")
	nums := []string{
		"aaabc89232",
		"10000",
		"13800138000",
	}

	for _, v := range nums {
		valid := validate.ValidMobilePhoneNumInChina(v)
		log.Printf("%v: %v", v, valid)
	}

	log.Printf("---------- ValidMobilePhoneNumInChina() Test End -------------")
	// Output:
}
