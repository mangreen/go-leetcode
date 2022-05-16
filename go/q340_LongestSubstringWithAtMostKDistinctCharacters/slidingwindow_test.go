package q340_LongestSubstringWithAtMostKDistinctCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slidewindow_output3(t *testing.T) {
	assert := assert.New(t)

	input := "eceba"
	ans := lengthOfLongestSubstringKDistinct(input, 2)

	assert.Equal(3, ans, "Longest substr \"ece\" length is 3.")
}

func Test_slidewindow_output2(t *testing.T) {
	assert := assert.New(t)

	input := "aa"

	ans := lengthOfLongestSubstringKDistinct(input, 1) 
	
	assert.Equal(2, ans, "Longest substr \"aa\" length is 5.")
}

func Test_slidewindow_output4(t *testing.T) {
	assert := assert.New(t)

	input := "aaccbc"

	ans := lengthOfLongestSubstringKDistinct(input, 2) 
	
	assert.Equal(4, ans, "Longest substr \"ccbc\" length is 4.")
}
