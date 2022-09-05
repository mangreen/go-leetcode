package q711_NumberOfDistinctIslandsII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_island_1(t *testing.T) {
	assert := assert.New(t)

	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 1},
	}
	ans := numDistinctIslands2(grid)

	assert.Equal(1, ans, " only 1 island")
}

func Test_islands_2(t *testing.T) {
	assert := assert.New(t)

	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 1, 1, 0, 0},
	}
	ans := numDistinctIslands2(grid)

	assert.Equal(2, ans, "has 2 islands")
}