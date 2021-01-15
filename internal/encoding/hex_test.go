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
