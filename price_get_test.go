package main

import "testing"

func TestPriceGet(t *testing.T) {
	dep := "RNO"
	arv := "SFO"

	got, err := getPrice(dep, arv)

	if err != nil || got <= 0.0 {
		t.Errorf("expected a valid price, got %v\nWith error: %v", got, err.Error())
	}
}
