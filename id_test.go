package validate_test

import (
	"log"

	"github.com/northbright/validate"
)

func ExampleValidIDCardNo() {
	log.Printf("---------- ValidIDCardNo() Test Begin -------------")
	nums := []string{
		"31010419810101400X",
		"310104199001013001",
		"310104600101001",
	}

	for _, v := range nums {
		valid := validate.ValidIDCardNo(v)
		log.Printf("%v: %v", v, valid)
	}

	log.Printf("---------- ValidIDCardNo() Test End -------------")
	// Output:
}
