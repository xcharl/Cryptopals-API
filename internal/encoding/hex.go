package algorithms

import (
	"errors"
	"strings"
)

const hexValues = "0123456789abcdef"

func Encode(input []byte) (string, error) {
	var output string

	for i := 0; i < len(input); i += 1 {
		mostSigBits := (input[i] & byte(0xf0)) >> 4
		leastSigBits := input[i] & byte(0xf)

		hex1 := string(rune(hexValues[mostSigBits]))
		hex2 := string(rune(hexValues[leastSigBits]))
		output += hex1 + hex2
	}

	return output, nil
}

func Decode(input string) ([]byte, error) {
	if len(input)%2 != 0 {
		return nil, errors.New("Cannot decode odd number of hex values.")
	}

	input = strings.ToLower(input)

	for i := 0; i < len(input); i += 2 {
		// mostSigBits := input[i]
		// leastSigBits := input[i + 1]

	}

	return nil, nil
}

func getByteFromHexChar(hexChar byte) byte {
	if hexChar >= 0x30 && hexChar <= 0x39 {
		// Numbers 0-9
		return hexChar - byte(0x30)
	} else if hexChar >= 0x61 && hexChar <= 0x66 {
		// Letters a-f
		return hexChar - byte(0x61) + 10
	}

	return byte(0)
}
