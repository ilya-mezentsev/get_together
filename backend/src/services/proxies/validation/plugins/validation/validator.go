package validation

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"regexp"
)

const (
	nicknamePattern = `^[-a-zA-Z0-9_$]{3,16}$`
	passwordPattern = `^[-a-zA-Z0-9_$*%&]{8,32}$`
	namePattern     = `^[a-zA-Za-zA-Zа-яА-ЯёЁ0-9. ]{3,16}$`
	textPattern     = `^[-_%$?:.,*()a-zA-Zа-яА-ЯёЁ0-9 ]+$`

	shortTextMinLength, shortTextMaxLength = 3, 255
	longTextMinLength, longTextMaxLength   = 15, 1024
	DateFormat                             = `02-01-2006 15:04:05`
)

var (
	textReg     = regexp.MustCompile(textPattern)
	nameReg     = regexp.MustCompile(namePattern)
	nicknameReg = regexp.MustCompile(nicknamePattern)
	passwordReg = regexp.MustCompile(passwordPattern)
)

func ValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func ValidNickname(n string) bool {
	return nicknameReg.MatchString(n)
}

func ValidLongitude(l float64) bool {
	return govalidator.IsLongitude(float64ToString(l))
}

func ValidLatitude(l float64) bool {
	return govalidator.IsLatitude(float64ToString(l))
}

func ValidWholePositiveNumber(n float64) bool {
	return govalidator.IsNatural(n)
}

func ValidPassword(p string) bool {
	return passwordReg.MatchString(p)
}

func ValidTitle(t string) bool {
	return textReg.MatchString(trim(t)) && govalidator.InRange(len(t), shortTextMinLength, shortTextMaxLength)
}

func ValidDescription(d string) bool {
	return textReg.MatchString(trim(d)) && govalidator.InRange(len(d), longTextMinLength, longTextMaxLength)
}

func ValidName(n string) bool {
	return nameReg.MatchString(trim(n))
}

func ValidDate(d string) bool {
	return govalidator.IsTime(d, DateFormat)
}

func ValidGender(g string) bool {
	return g == "male" || g == "female" || g == ""
}

func ValidURL(u string) bool {
	return govalidator.IsURL(u)
}

func trim(s string) string {
	return govalidator.Trim(s, "")
}

func float64ToString(f float64) string {
	return fmt.Sprintf("%f", f)
}
