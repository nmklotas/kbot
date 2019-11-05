package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsesListOrders(t *testing.T) {
	assert := assert.New(t)
	result := ParseListOrders("kbot list")
	assert.True(result, "kbot list not parsed")
}

func TestParsesOrderAdd(t *testing.T) {
	assert := assert.New(t)
	result, _ := ParseOrderToAdd("kbot order A")
	assert.Equal("A", *result, "kbot list not parsed")
}

func TestParsesOrderRemove(t *testing.T) {
	assert := assert.New(t)
	result, _ := ParseOrderToRemove("kbot forget A")
	assert.Equal("A", *result, "kbot list not parsed")
}
