package encoding

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type encodeTestCase struct {
	input []byte
	expected string
	errExpected bool
}

func TestEncodeHex(t *testing.T) {
	encodeTestCases := []encodeTestCase {
		{ nil, "", true, },
		{ []byte{}, "", false, },
		{ []byte{0xff}, "ff", false, },
		{ []byte{00}, "00", false, },
		{ []byte{0x26, 0xab, 0x36, 0xff}, "26ab36ff", false, },
	}

	for _, test := range encodeTestCases {
		actual, err := EncodeHex(test.input)

		assert.Equal(t, test.expected, actual)

		if (test.errExpected) {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

type decodeTestCase struct {
	input string
	expected []byte
	errExpected bool
}

func TestDecodeHex(t *testing.T) {
	decodeTestCases := []decodeTestCase {
		{ "", []byte{}, false, },
		{ "",[]byte{},  false, },
		{  "ff",[]byte{0xff}, false, },
		{ "00", []byte{00}, false, },
		{ "26ab36ff", []byte{0x26, 0xab, 0x36, 0xff}, false, },
		{ "26AB36FF", []byte{0x26, 0xab, 0x36, 0xff}, false, },
		{ "42helloe", nil, true, }, // Invalid hex chars
		{ "123456789", nil, true, }, // Invalid input length
		{ "abcdefghi", nil, true, }, // Invalid len and hex char
	}

	for _, test := range decodeTestCases {
		actual, err := DecodeHex(test.input)

		assert.Equal(t, test.expected, actual)

		if (test.errExpected) {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

