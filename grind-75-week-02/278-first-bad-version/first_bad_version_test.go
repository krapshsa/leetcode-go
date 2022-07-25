package first_bad_version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstBadVersion_4_5(t *testing.T) {
	badVersion = 4
	assert.Equal(t, badVersion, firstBadVersion(5))
}

func Test_firstBadVersion_1_1(t *testing.T) {
	badVersion = 1
	assert.Equal(t, badVersion, firstBadVersion(1))
}

func Test_firstBadVersion_1_3(t *testing.T) {
	badVersion = 1
	assert.Equal(t, badVersion, firstBadVersion(3))
}

func Test_firstBadVersion_2_4(t *testing.T) {
	badVersion = 2
	assert.Equal(t, badVersion, firstBadVersion(4))
}

func Test_firstBadVersion_2_2(t *testing.T) {
	badVersion = 2
	assert.Equal(t, badVersion, firstBadVersion(2))
}
