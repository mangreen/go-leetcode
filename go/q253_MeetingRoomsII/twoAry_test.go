package q252_MeetingRooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMeetingRooms_false(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{0,30}, {5,10}, {15,20}}
	ans := minMeetingRooms(input)

	assert.Equal(2, ans, "[[0,30],[5,10],[15,20]] has overlap [0,30] & [5,10], so needs 2 rooms.")
}

func Test_minMeetingRooms_true(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{7,10}, {2,4}}
	ans := minMeetingRooms(input)

	assert.Equal(1, ans, "[[7,10],[2,4]] no overlap, so only need 1 room.")
}
