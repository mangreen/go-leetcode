package q261_GraphValidTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unionfind_isTrue(t *testing.T) {
	assert := assert.New(t)

	ans := validTree2(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}})

	assert.Equal(true, ans, "true, It's a tree")
}

func Test_unionfind_isFalse(t *testing.T) {
	assert := assert.New(t)

	ans := validTree2(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}})

	assert.Equal(false, ans, "false, It's not a tree")
}

func Test_unionfind_isFalse2(t *testing.T) {
	assert := assert.New(t)

	ans := validTree2(5, [][]int{{0, 1}, {2, 3}, {1, 4}})

	assert.Equal(false, ans, "false, It's not a tree")
}
