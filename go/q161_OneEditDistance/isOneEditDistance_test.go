package q161_OneEditDistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_true(t *testing.T) {
	assert := assert.New(t)

	ans := isOneEditDistance("ab", "acb")

	assert.Equal(true, ans, "ab can OneEditDistance to be acb")
}

func Test_false(t *testing.T) {
	assert := assert.New(t)

	ans := isOneEditDistance("cab", "ad")

	assert.Equal(false, ans, "cab can't OneEditDistance to be ad")
}