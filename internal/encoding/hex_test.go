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
		encodeTestCase { nil, "", true, },
		encodeTestCase { []byte{}, "", false, },
		encodeTestCase { []byte{0xff}, "ff", false, },
		encodeTestCase { []byte{00}, "00", false, },
		encodeTestCase { []byte{0x26, 0xab, 0x36, 0xff}, "26ab36ff", false, },
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
		decodeTestCase { "", []byte{}, false, },
		decodeTestCase { "",[]byte{},  false, },
		decodeTestCase {  "ff",[]byte{0xff}, false, },
		decodeTestCase { "00", []byte{00}, false, },
		decodeTestCase { "26ab36ff", []byte{0x26, 0xab, 0x36, 0xff}, false, },
		decodeTestCase { "26AB36FF", []byte{0x26, 0xab, 0x36, 0xff}, false, },
		decodeTestCase { "42helloe", nil, true, }, // Invalid hex chars
		decodeTestCase { "123456789", nil, true, }, // Invalid input length
		decodeTestCase { "abcdefghi", nil, true, }, // Invalid len and hex char
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

