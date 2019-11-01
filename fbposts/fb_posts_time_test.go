package fbposts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
