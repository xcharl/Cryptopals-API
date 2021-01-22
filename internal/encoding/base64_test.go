package encoding

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestEncodeStringToBase64(t *testing.T) {
	input := "AAA"
	expected := "QUFB"

	actual := EncodeStringToBase64(input)

	assert.Equal(t, expected, actual)
}

