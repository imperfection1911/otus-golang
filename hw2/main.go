package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Unpack function SUDDENLY returns unpacked string
func Unpack(input string) (string, error) {
	var previous rune
	var unpacked string
	var builder strings.Builder
	for _, char := range input {
		if unicode.IsNumber(char) {
			if unicode.IsNumber(previous) {
				return "", errors.New("invalid string")
			}
			// мне не нравится этот способ конвертировать руну в инт. Есть альтернатива?
			count := int(char - '0')
			for i := 0; i < count-1; i++ {
				builder.WriteRune(previous)
			}
		} else {
			builder.WriteRune(char)
		}
		previous = char
	}
	unpacked = builder.String()
	return unpacked, nil
}

func main() {
	result, unpackError := Unpack("q5wer3t7y1")
	if unpackError != nil {
		fmt.Println(unpackError)
		os.Exit(1)
	}
	fmt.Println(result)

}
