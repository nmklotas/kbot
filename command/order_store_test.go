package command

import "testing"

func TestCanAdd(t *testing.T) {
	db, err := NewOrdersStore()
	if err != nil {
		t.Errorf("store not created")
	}

	order := "A"
	err = db.Add("1", &order)
	if err != nil {
		t.Errorf("order not added")
	}

	orders, err := db.List()
	if err != nil || len(orders) == 0 {
		t.Errorf("order not added")
	}
}

func TestCanRemove(t *testing.T) {
	db, err := NewOrdersStore()
	if err != nil {
		t.Errorf("store not created")
	}

	order := "A"
	err = db.Add("1", &order)
	if err != nil {
		t.Errorf("order not added")
	}

	err = db.Remove("1", &order)
	if err != nil {
		t.Errorf("order not removed")
	}

	orders, err := db.List()
	if err != nil || len(orders) != 0 {
		t.Errorf("order not removed")
	}
}
