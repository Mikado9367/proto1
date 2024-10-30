package tool

import (
	"regexp"
	"strconv"
	"time"
)

func IsDateValue(stringDate string) bool {

	if _, err := time.Parse("2006-01-02", stringDate); err != nil {
		return false
	} else {
		return true
	}
}

func IsAlphaNumeric(s string) bool {

	return regexp.MustCompile(`^[a-zA-Z0-9]`).MatchString(s)
}

func StringToInt(str string) (i int, err error) {

	n, err := strconv.Atoi(str)

	return n, err
}

func IntToString(i int) (str string) {

	str = strconv.Itoa(i)

	return str
}
