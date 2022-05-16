package q266_PermutePalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isTrue(t *testing.T) {
	assert := assert.New(t)

	ans := canPermutePalindrome("tactcoa")

	assert.Equal(true, ans, "they should be equal")
}

func Test_isFalse(t *testing.T) {
	assert := assert.New(t)

	ans := canPermutePalindrome("abc")

	assert.Equal(false, ans, "they should be equal")
}
