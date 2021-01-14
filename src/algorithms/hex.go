package algorithms

import (
	"errors"
)

const hexValues = "0123456789abcdef"

func EncodeToHex(input []byte) (string, error) {
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

func DecodeFromHex(input string) ([]byte, error) {
	if len(input)%2 != 0 {
		return nil, errors.New("Cannot decode odd number of hex values.")
	}

	for i := 0; i < len(input); i += 2 {
		// mostSigBits := input[i]
		// leastSigBits := input[i + 1]

	}

	return nil, nil
}
