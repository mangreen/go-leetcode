package q359_LoggerRateLimiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shouldPrintMessage(t *testing.T) {
	assert := assert.New(t)

	log := Constructor()

	assert.Equal(log.shouldPrintMessage(1, "foo"), true, "It's true")
	assert.Equal(log.shouldPrintMessage(2, "bar"), true, "It's true")
	assert.Equal(log.shouldPrintMessage(3, "foo"), false, "It's false")
	assert.Equal(log.shouldPrintMessage(8, "bar"), false, "It's false")
	assert.Equal(log.shouldPrintMessage(10, "foo"), false, "It's false")
	assert.Equal(log.shouldPrintMessage(111, "foo"), true, "It's true")
}
