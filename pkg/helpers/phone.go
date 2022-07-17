package helpers

import (
	"fmt"
	"strings"
)

const (
	VNPhoneCode = "84"
)

func SanitizePhoneNumber(phoneCode, phoneNumber string) string {
	phoneCode = strings.TrimLeft(phoneCode, "+")
	if phoneCode == VNPhoneCode {
		phoneNumber = strings.TrimLeft(phoneNumber, "0")
	}

	return fmt.Sprintf("%s%s", phoneCode, phoneNumber)
}
