package encoding

import "fmt"

const base64Lookup = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const padding = "="

// EncodeStringToBase64 encodes a string to Base64
func EncodeStringToBase64(input string) string {
	return EncodeBase64([]byte(input))
}

// EncodeBase64 encodes a byte array to Base64
func EncodeBase64(input []byte) string {
	var outputLen int
	if len(input) % 3 == 0 {
		outputLen = len(input) * 4 / 3
	} else {
		outputLen = (len(input) + 4) * 4 / 3
	}

	output := make([]byte, outputLen)

	// Process in blocks of 3 bytes (== 4 sextets out)
	var inInd, outInd int
	for ; inInd < len(input); inInd += 3 {
		out1 := input[inInd] >> 2
		out2 := (input[inInd] & 0b00000011 << 4) | (input[inInd+1] >> 4)
		out3 := (input[inInd+1] & 0b00001111 << 2) | (input[inInd+2] >> 6)
		out4 := input[inInd+2] & 0b00111111

		output[outInd] = base64Lookup[out1]
		output[outInd+1] = base64Lookup[out2]
		output[outInd+2] = base64Lookup[out3]
		output[outInd+3] = base64Lookup[out4]

		fmt.Println(string(out1), (out2), out3, (outInd), out4)
	}

	return string(output)
}
