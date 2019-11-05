package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAdd(t *testing.T) {
	assert := assert.New(t)
	db, err := OpenOrdersRepository()
	if err != nil {
		t.Errorf("store not created")
	}

	err = db.Add(Order{UserId: "1", Value: "A"})
	if err != nil {
		t.Errorf("order not added")
	}

	orders, _ := db.List()
	assert.NotEmpty(orders)
}

func TestCanRemove(t *testing.T) {
	assert := assert.New(t)
	db, err := OpenOrdersRepository()
	if err != nil {
		t.Errorf("store not created")
	}

	err = db.Add(Order{UserId: "1", Value: "A"})
	if err != nil {
		t.Errorf("order not added")
	}

	err = db.Remove("1", "A")
	if err != nil {
		t.Errorf("order not removed")
	}

	orders, _ := db.List()
	assert.Empty(orders)
}
