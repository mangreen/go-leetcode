package q683_KEmptySlots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kEmptySlots_1(t *testing.T) {
	assert := assert.New(t)

	bulbs := []int{1, 3, 2}
	k := 1
	ans := kEmptySlots(bulbs, k)

	assert.Equal(2, ans, "2nd day is the answer")
}

func Test_kEmptySlots_2(t *testing.T) {
	assert := assert.New(t)

	bulbs := []int{1, 2, 3}
	k := 1
	ans := kEmptySlots(bulbs, k)

	assert.Equal(-1, ans, "Can't find the day")
}