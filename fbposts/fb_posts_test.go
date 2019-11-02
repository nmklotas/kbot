package fbposts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsWord(t *testing.T) {
	assert := assert.New(t)
	assert.True(ContainsWord(FbPost{Text: "word123"}, "word"))
}

func TestNotContainsWord(t *testing.T) {
	assert := assert.New(t)
	assert.False(ContainsWord(FbPost{Text: "wrdo123"}, "word"))
}
