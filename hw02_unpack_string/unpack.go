package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		builder strings.Builder
		runes   []rune
	)

	runes = []rune(s)
	if len(runes) == 0 {
		return "", nil
	}
	for i := range runes {
		if unicode.IsDigit(runes[i]) {
			if i == 0 || unicode.IsDigit(runes[i-1]) {
				return "", ErrInvalidString
			}

			digit, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", ErrInvalidString
			}

			if digit == 0 {
				continue
			}
			builder.WriteString(strings.Repeat(string(runes[i-1]), digit))
		}

		if !unicode.IsDigit(runes[i]) && isLiteralSingle(i, runes) {
			builder.WriteRune(runes[i])
		}
	}
	return builder.String(), nil
}

func isLiteralSingle(i int, runes []rune) bool {
	return i == len(runes)-1 || !unicode.IsDigit(runes[i+1]) &&
		(i == 0 || i != len(runes)-1)
}
