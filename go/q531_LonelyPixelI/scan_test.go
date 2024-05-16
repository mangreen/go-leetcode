package q683_KEmptySlots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLonelyPixel(t *testing.T) {
	assert := assert.New(t)

	picture := [][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'},
	}

	ans := findLonelyPixel(picture)

	assert.Equal(3, ans, "There are 3 lonely pixel")
}
