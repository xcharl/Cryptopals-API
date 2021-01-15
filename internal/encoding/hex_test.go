package encoding

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// type encodeHexTestCase struct {
// 	input []byte
// 	expected string
// 	output string
// 	errExpected bool
// }

func TestEncodeHexNilInputReturnsEmptyString(t *testing.T) {
	expected := ""

	var input []byte
	actual, err := EncodeHex(input)
	
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeHexFf(t *testing.T) {
	expected := "ff"

	input := []byte{0xff}
	actual, err := EncodeHex(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeHex00(t *testing.T) {
	expected := "00"

	input := []byte{0x00}
	actual, err := EncodeHex(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeHexRandomBytes(t *testing.T) {
	expected := "26ab36ff"

	input := []byte{0x26, 0xab, 0x36, 0xff}
	actual, err := EncodeHex(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
