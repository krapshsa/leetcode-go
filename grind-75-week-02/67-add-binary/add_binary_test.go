package add_binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "10101", addBinary("1010", "1011"))
	assert.Equal(t, "0", addBinary("0", "0"))
}
