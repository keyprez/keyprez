package utils

import (
	"net/mail"

	"github.com/go-playground/validator"
)

func IsValidEmail(email string) bool {
	_, emailErr := mail.ParseAddress(email)
	return !(emailErr != nil || email == "")
}

func IsValidCustomer(requestBody CreateCustomerRequest) bool {
	validate := validator.New()
	err := validate.Struct(requestBody)
	return err == nil
}
