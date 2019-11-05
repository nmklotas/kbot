package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFindWithTruePredicate(t *testing.T) {
	assert := assert.New(t)
	db, err := OpenLunchRepository()
	if err != nil {
		t.Errorf("store not created")
	}

	err = db.Save(Lunch{})
	if err != nil {
		t.Errorf("store not created")
	}

	result := db.Any(func(l Lunch) bool {
		return true
	})

	assert.True(result, "order not found")
}

func TestCantFindWithFalsePredicate(t *testing.T) {
	assert := assert.New(t)
	db, err := OpenLunchRepository()
	if err != nil {
		t.Errorf("store not created")
	}

	err = db.Save(Lunch{})
	if err != nil {
		t.Errorf("lunch not saved")
	}

	result := db.Any(func(l Lunch) bool {
		return false
	})

	assert.False(result, "order should not be found")
}
