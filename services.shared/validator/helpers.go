package validator

import "regexp"

func CheckEmail(email string) bool {
	matched, _ := regexp.MatchString(emailRegexp, email)
	return matched
}

func CheckPhone(phone string) bool {
	matched, _ := regexp.MatchString(PhoneWithCountryCodeRegexp, phone)
	return matched
}
