package check

import (
	"errors"
	"regexp"
)

// Employee check employee flag and return error if value don't match regexp
func Employee(value string) error {
	match, err := regexp.MatchString(`^[а-яА-Яa-zA-z][а-яa-z-А-ЯA-Z-_\.]{2,20}$`, value)
	if err != nil {
		return err
	} else if !match {
		return errors.New("invalid employee. allowed only latters")
	}

	return nil
}

// Date check date flags and return error if date don't match regexp DD.MM.YYYY
func Date(date string) error {
	match, err := regexp.MatchString(`(0[1-9]|[12][0-9]|3[01])[- ..](0[1-9]|1[012])[- ..][201]\d\d\d`, date)
	if err != nil {
		return err
	} else if !match {
		return errors.New("invalid date format. correct format: DD.MM.YYYY")
	}

	return nil
}
