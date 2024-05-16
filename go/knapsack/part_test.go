package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var itemsWithCount = [][]int{
	{5, 8, 2}, 
    {4, 3, 2},
    {7, 10, 2}, 
    {2, 3, 3},
    {6, 6, 3}, 
}

func Test_part_15(t *testing.T) {
	assert := assert.New(t)

	ans := complete(itemsWithCount, 15)

	assert.Equal(22, ans, "The value of W=15 is 22")
}

func Test_part_5(t *testing.T) {
	assert := assert.New(t)

	ans := complete(itemsWithCount, 5)

	assert.Equal(8, ans, "The value of W=15 is 8")
}