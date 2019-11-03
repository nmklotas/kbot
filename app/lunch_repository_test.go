package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFindWithTruePredicate(t *testing.T) {
	assert := assert.New(t)
	db, err := OpenLunchRepository()
	assert.NoError(err)
	err = db.Save(Lunch{})
	assert.NoError(err)

	result := db.Any(func(l Lunch) bool {
		return true
	})

	assert.True(result, "order not found")
}

func TestCantFindWithFalsePredicate(t *testing.T) {
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

	if result {
		t.Errorf("order should not be found")
	}
}
