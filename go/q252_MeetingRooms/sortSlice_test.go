package q252_MeetingRooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canAttendMeetings_false(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{0,30}, {5,10}, {15,20}}
	ans := canAttendMeetings(input)

	assert.Equal(false, ans, "[[0,30],[5,10],[15,20]] has overlap [0,30] & [5,10].")
}

func Test_canAttendMeetings_true(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{7,10}, {2,4}}
	ans := canAttendMeetings(input)

	assert.Equal(true, ans, "[[7,10],[2,4]] doesn't have overlap.")
}

func Test_canAttendMeetings_false2(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{{1,2}, {1,3}, {2,3}, {3,4}}
	ans := canAttendMeetings(input)

	assert.Equal(false, ans, "[[1,2], [1,3], [2,3], [3,4]] has overlap [1, 3].")
}
