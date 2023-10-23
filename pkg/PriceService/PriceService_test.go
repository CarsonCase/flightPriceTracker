package PriceService

import (
	"testing"
)

func TestPriceGet(t *testing.T) {
	t.Run("getPrice returns non-zero value", func(t *testing.T) {
		dep := "RNO"
		arv := "SFO"

		got, err := GetPrice(dep, arv, "2024-01-01")

		if err != nil || got <= 0.0 {
			t.Errorf("expected a valid price, got %v\nWith error: %v", got, err.Error())
		}
	})
}
