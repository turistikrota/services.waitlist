package validator

import (
	"regexp"

	"opensource.turistikrota.com/shared/helper"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateUserName(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(userNameRegexp, fl.Field().String())
	return matched
}

func validatePassword(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(passwordRegexp, fl.Field().String())
	return matched
}

func validateSlug(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(slugRegexp, fl.Field().String())
	return matched
}

func validateLocale(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(localeRegexp, fl.Field().String())
	return matched
}

func validateObjectId(fl validator.FieldLevel) bool {
	return primitive.IsValidObjectID(fl.Field().String())
}

func validateGender(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(genderRegexp, fl.Field().String())
	return matched
}

func validatePhone(f1 validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(PhoneWithCountryCodeRegexp, f1.Field().String())
	return matched
}

func validateUUID(f1 validator.FieldLevel) bool {
	_, err := uuid.Parse(f1.Field().String())
	return err == nil
}

func validateUserNameAndCode(f1 validator.FieldLevel) bool {
	val := f1.Field().String()
	matched, _ := regexp.MatchString(userNameAndCodeRegexp, val)
	if !matched {
		return false
	}
	matched, _, _ = helper.Parsers.ParseUsernameAndCode(val)
	return matched
}
