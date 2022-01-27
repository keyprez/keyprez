package utils

import (
	"net/mail"
)

func IsValidEmail(email string) bool {
	_, emailErr := mail.ParseAddress(email)
	return !(emailErr != nil || email == "")
}
