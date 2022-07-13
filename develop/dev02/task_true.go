package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Enum struct {
	Value int
}

func (e *Enum) CheckValue() string {
	if e.Value == 1 {
		return "BEGIN"
	} else if e.Value == 2 {
		return "NUMBER"
	} else if e.Value == 3 {
		return "SYMBOL"
	} else if e.Value == 4 {
		return "ESCAPE"
	} else {
		return "ERROR"
	}
}

func CheckEscape(i rune) bool {
	return i == '\\'
}

func unpackString(input string) (string, error) {
	var builder strings.Builder
	if len(input) == 0 {
		return "", errors.New("empty string")
	}
	var enum Enum
	enum.Value = 1
	for i := 0; i < len(input); i++ {
		//fmt.Println(string(input[i])) //BEGIN - 1; NUMBER - 2; SYMBOL - 3; ESCAPE - 4
		if enum.CheckValue() == "BEGIN" && unicode.IsDigit(rune(input[i])) {
			return "", errors.New("first number")
		} else if enum.CheckValue() == "BEGIN" && unicode.IsLetter(rune(input[i])) {
			builder.WriteRune(rune(input[i]))
			enum.Value = 3
		} else if enum.CheckValue() == "SYMBOL" && unicode.IsLetter(rune(input[i])) {
			builder.WriteRune(rune(input[i]))
			enum.Value = 3
		} else if enum.CheckValue() == "SYMBOL" && unicode.IsDigit(rune(input[i])) {
			number, err := strconv.Atoi(string(input[i]))
			if err != nil {
				return "", errors.New("error in number")
			}
			if number >= 0 {
				for j := 0; j < number-1; j++ {
					builder.WriteRune(rune(input[i-1]))
				}
			}
			enum.Value = 2
		} else if enum.CheckValue() == "SYMBOL" && unicode.IsLetter(rune(input[i])) {
			builder.WriteRune(rune(input[i]))
			enum.Value = 3
		} else if enum.CheckValue() == "SYMBOL" && CheckEscape(rune(input[i])) {
			enum.Value = 4
		} else if enum.CheckValue() == "NUMBER" && unicode.IsDigit(rune(input[i])) {
			return input, errors.New("error in number")
		} else if enum.CheckValue() == "NUMBER" && unicode.IsLetter(rune(input[i])) {
			builder.WriteRune(rune(input[i]))
			enum.Value = 3
		} else if enum.CheckValue() != "ESCAPE" && CheckEscape(rune(input[i])) {
			enum.Value = 4
		} else if enum.CheckValue() == "ESCAPE" && !CheckEscape(rune(input[i])) {
			builder.WriteRune(rune(input[i]))
			enum.Value = 3
		} else if enum.CheckValue() == "ESCAPE" && CheckEscape(rune(input[i])) {
			builder.WriteRune(rune(input[i-1]))
			enum.Value = 3
		}
	}
	if enum.CheckValue() == "ESCAPE" {
		return "", errors.New("escape sequence")
	}
	return builder.String(), nil
}

func main() {
}
