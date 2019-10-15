package command

import "testing"

func TestParsesListOrders(t *testing.T) {
	result := ParseListOrders("kbot list")
	if !result {
		t.Errorf("kbot list not parsed")
	}
}

func TestParsesOrderAdd(t *testing.T) {
	result, err := ParseOrderToAdd("kbot order A")
	if err != nil {
		t.Errorf("kbot order A parse error")
	}

	if *result != "A" {
		t.Errorf("kbot order A not parsed")
	}
}

func TestParsesOrderRemove(t *testing.T) {
	result, err := ParseOrderToRemove("kbot forget A")
	if err != nil {
		t.Errorf("kbot forget A parse error")
	}

	if *result != "A" {
		t.Errorf("kbot forget A not parsed")
	}
}
