package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var items = [][]int{
	{5, 12}, 
	{4, 3},
	{7, 10}, 
	{2, 3},
	{6, 6}, 
}

func Test_complete_15(t *testing.T) {
	assert := assert.New(t)

	ans := complete(items, 15)

	assert.Equal(36, ans, "The value of W=15 is 36")
}

func Test_complete_5(t *testing.T) {
	assert := assert.New(t)

	ans := complete(items, 5)

	assert.Equal(12, ans, "The value of W=15 is 12")
}