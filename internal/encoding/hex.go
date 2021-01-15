package encoding

import (
	"errors"
	"strings"
)

const hexValues = "0123456789abcdef"

func EncodeHex(input []byte) (string, error) {
	var output string

	if input == nil {
		return output, errors.New("Nil input.")
	}

	for i := 0; i < len(input); i += 1 {
		mostSigBits := (input[i] & byte(0xf0)) >> 4
		leastSigBits := input[i] & byte(0xf)

		hex1 := string(rune(hexValues[mostSigBits]))
		hex2 := string(rune(hexValues[leastSigBits]))
		output += hex1 + hex2
	}

	return output, nil
}

func DecodeHex(input string) ([]byte, error) {
	if len(input)%2 != 0 {
		return nil, errors.New("Cannot decode odd number of hex values.")
	}

	input = strings.ToLower(input)
	output := make([]byte, len(input)/2)

	for i := 0; i < len(input); i += 2 {
		mostSigBits, err1 := getBitsFromHex(input[i])
		leastSigBits, err2 := getBitsFromHex(input[i + 1])

		if err1 != nil {
			return nil, err1
		}

		if err2 != nil {
			return nil, err2
		}

		currByte := (mostSigBits << 4) | leastSigBits
		output[i/2] = currByte
	}

	return output, nil
}

func getBitsFromHex(hexChar byte) (byte, error) {
	if hexChar >= 0x30 && hexChar <= 0x39 {
		// Numbers 0-9 in ASCII
		return hexChar - byte(0x30), nil
	} else if hexChar >= 0x61 && hexChar <= 0x66 {
		// Letters a-f
		return hexChar - byte(0x61) + 10, nil
	}

	return 0, errors.New("Invalid hex character encountered: " + string(hexChar))
}
