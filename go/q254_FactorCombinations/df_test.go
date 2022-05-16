package q254_FactorCombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPath(t *testing.T) {
	assert := assert.New(t)

	ans := getFactors(1)

	expect := [][]int{}

	assert.Equal(expect, ans, "[]")
}

func Test_hasPath_failed(t *testing.T) {
	assert := assert.New(t)

	ans := getFactors(12)

	expect := [][]int{{2, 2, 3}, {2, 6}, {3, 4}}

	assert.Equal(expect, ans, "[[2, 6], [2, 2, 3], [3, 4]]")
}