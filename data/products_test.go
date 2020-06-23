package data

import "testing"

// TestCheckValidation is a simple dummy unit test
func TestCheckValidation(t *testing.T)  {
	p := &Product{
		Name:"naren",
		Price: 100.00,
		SKU: "aas-fes-rwg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}