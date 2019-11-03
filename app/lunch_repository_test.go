package app

import "testing"

func TestCanAdd(t *testing.T) {
	db, err := OpenLunchRepository()
	if err != nil {
		t.Errorf("store not created")
	}

	err = db.Save(Lunch{})
	if err != nil {
		t.Errorf("lunch not saved")
	}

	orders, err := db.List()
	if err != nil || len(orders) == 0 {
		t.Errorf("order not added")
	}
}
