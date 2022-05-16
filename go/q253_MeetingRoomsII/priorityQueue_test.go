package q252_MeetingRooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMeetingRooms2_false(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{0,30}, {5,10}, {15,20}}
	ans := minMeetingRooms2(input)

	assert.Equal(2, ans, "[[0,30],[5,10],[15,20]] needs 2 rooms.")
}

func Test_minMeetingRooms2_true(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{7,10}, {2,4}}
	ans := minMeetingRooms2(input)

	assert.Equal(1, ans, "[[7,10],[2,4]] no overlap, so only need 1 room.")
}
