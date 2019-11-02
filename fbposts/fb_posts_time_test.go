package fbposts

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsTimeToCheckForPost(t *testing.T) {
	assert := assert.New(t)

	createTime := func(hour int) time.Time {
		time, _ := time.Parse(time.RFC3339, fmt.Sprintf("2014-11-12T%d:45:26.371Z", hour))
		return time
	}

	checkInterval := CheckInterval{4 * 60, 5 * 60}
	assert.True(IsTimeToCheck(createTime(7), createTime(3), checkInterval))
	assert.False(IsTimeToCheck(createTime(7), createTime(1), checkInterval))
	assert.False(IsTimeToCheck(createTime(7), createTime(4), checkInterval))
}

func TestCanParseFacebookTimeToLocal(t *testing.T) {
	assert := assert.New(t)

	_, err := ParseToLocalTime("2019-11-01T21:47:06+0000")
	assert.NoError(err)
}
