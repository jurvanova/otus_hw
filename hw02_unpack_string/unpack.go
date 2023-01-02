package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	ru := []rune(s)
	if len(ru) == 0 {
		return "", nil
	}
	var newStr strings.Builder
	for i := range ru {
		if unicode.IsDigit(ru[i]) && (i == 0 || unicode.IsDigit(ru[i-1])) {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(ru[i]) && (i == len(ru)-1 || (i == 0 && !unicode.IsDigit(ru[i+1])) ||
			(i != len(ru)-1 && unicode.IsLetter(ru[i+1]))) {
			newStr.WriteRune(ru[i])
			continue
		}

		if unicode.IsDigit(ru[i]) && (unicode.IsLetter(ru[i-1]) ||
			unicode.IsSpace(ru[i-1])) {
			digit, err := strconv.Atoi(string(ru[i]))
			if err != nil {
				return "", ErrInvalidString
			}

			if digit == 0 {
				continue
			}

			newStr.WriteString(strings.Repeat(string(ru[i-1]), digit))
			continue
		}

	}
	return newStr.String(), nil
}
