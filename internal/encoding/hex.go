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
	var output []byte

	for i := 0; i < len(input); i += 2 {
		mostSigBits := input[i] << 4
		leastSigBits := input[i + 1]
		currByte := mostSigBits | leastSigBits
		output[i/2] = currByte
	}

	return output, nil
}

func getByteFromHexChar(hexChar byte) (byte, error) {
	if hexChar >= 0x30 && hexChar <= 0x39 {
		// Numbers 0-9
		return hexChar - byte(0x30), nil
	} else if hexChar >= 0x61 && hexChar <= 0x66 {
		// Letters a-f
		return hexChar - byte(0x61) + 10, nil
	}

	return 0, errors.New("Hex character not recognised.")
}
