package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Lemon Tea",
		Price: 1.00,
		SKU:   "abc-acd-bal",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
