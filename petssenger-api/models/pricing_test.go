package models

import (
	"testing"

	. "./pricing"
)

func TestCalculatorPricing(t *testing.T) {
	p := Pricing{
		City:           "São Paulo",
		BaseFee:        3.50,
		PricePerMinute: 1.00,
		ServiceFee:     0.75,
		PricePerKm:     0.50,
	}
	total := p.Calc(3, 10, 1)

	if total != 15.75 {
		t.Errorf("Sum was incorrect, got: %b, want: %b.", total, 15.75)
	}

}
