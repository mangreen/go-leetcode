package q159_LongestSubstringWithAtMost2DistinctCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slidingWindow_output3(t *testing.T) {
	assert := assert.New(t)

	input := "eceba"
	ans := lengthOfLongestSubstringTwoDistinct2(input)

	assert.Equal(3, ans, "Longest substr \"ece\" length is 3.")
}

func Test_slidingWindow_output5(t *testing.T) {
	assert := assert.New(t)

	input := "ccaabbb"

	ans := lengthOfLongestSubstringTwoDistinct2(input) 
	
	assert.Equal(5, ans, "Longest substr \"aabbb\" length is 5.")
}
