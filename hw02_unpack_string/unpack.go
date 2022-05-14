package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const winSlash = 92 // mean "\"

var (
	ErrInvalidString   = errors.New("invalid string")
	ErrInvalidEscaping = errors.New("slash before a wrong simbol")
)

func Unpack(str string) (string, error) {
	runes := []rune(str)
	var sb strings.Builder
	for i := 0; i < len(runes); {
		if unicode.IsDigit(runes[i]) {
			return "", ErrInvalidString
		}
		if runes[i] == winSlash {
			if i < len(runes)-1 && !(unicode.IsDigit(runes[i+1]) || runes[i+1] == winSlash) {
				return "", ErrInvalidEscaping
			}
			i++
		}
		if i < len(runes)-1 {
			if unicode.IsDigit(runes[i+1]) {
				num, _ := strconv.Atoi(string(runes[i+1]))
				ss := strings.Repeat(string(runes[i]), num)
				sb.WriteString(ss)
				i += 2
			} else {
				sb.WriteRune(runes[i])
				i++
			}
		} else {
			sb.WriteRune(runes[i])
			i++
		}
	}
	return sb.String(), nil
}
