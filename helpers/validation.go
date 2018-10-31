package helpers

import (
	"regexp"
)

// Employee validation employee flag
// return false if value don't match regexp
func ValidationEmployee(value string) bool {
	if match, _ := regexp.MatchString(
		`^[а-яА-Яa-zA-z][а-яa-z-А-ЯA-Z-_\.]{0,20}$`, value,
	); !match {
		return false
	}
	return true
}

// Date validation date flags
// return false if date don't match
// regexp DD.MM.YYYY or DD-MM-YYYY
func ValidationDate(date string) bool {
	if match, _ := regexp.MatchString(
		`(0[1-9]|[12][0-9]|3[01])[- ..](0[1-9]|1[012])[- ..][201]\d\d\d`, date,
	); !match {
		return false
	}
	return true
}
