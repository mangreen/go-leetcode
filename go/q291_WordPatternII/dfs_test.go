package q291_WordPatternII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPath(t *testing.T) {
	assert := assert.New(t)

	ans := wordPatternMatch("abab", "redblueredblue")

	assert.Equal(true, ans, "abab & redblueredblue have the same pattern")
}

func Test_hasPath_failed(t *testing.T) {
	assert := assert.New(t)

	ans := wordPatternMatch("aabb", "xyzabcxzyabc")

	assert.Equal(false, ans, "aabb & xyzabcxzyabc  not have the same pattern")
}