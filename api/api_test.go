package api

import (
	"testing"

	"github.com/drunknsorry/Tax-calculator/apiconsumer"
)

func TestCalculateTax(t *testing.T) {
	data := &apiconsumer.TaxBracketResults{
		TaxBrackets: []apiconsumer.Brackets{
			{Min: 0, Max: 50000, Rate: 0.10},
			{Min: 50000, Max: 75000, Rate: 0.05},
			{Min: 75000, Max: 100000, Rate: 0.20},
		},
	}

	year := 2019
	totalSalary := 100000

	tax, taxesPerBand, salaryPerBand, effectiveTaxRate := CalculateTaxes(data, float64(year), float64(totalSalary))

	expectedTaxesOwed := []float64{5000, 1250, 5000}
	for i := 0; i < len(taxesPerBand); i++ {
		if taxesPerBand[i] != expectedTaxesOwed[i] {
			t.Errorf("Expected Taxes Owed Per Band: %f, Got: %f", expectedTaxesOwed[i], taxesPerBand[i])
		}
	}

	expectedSalaryPerBand := []float64{50000, 25000, 25000}
	for i := 0; i < len(salaryPerBand); i++ {
		if salaryPerBand[i] != expectedSalaryPerBand[i] {
			t.Errorf("Expected Salary Amount Per Tax Band: %f, Got: %f", expectedSalaryPerBand[i], salaryPerBand[i])
		}
	}

	expectedEffectiveTaxRate := 0.11
	if expectedEffectiveTaxRate != effectiveTaxRate {
		t.Errorf("Expected Effective Tax Rate: %f, Got: %f", expectedEffectiveTaxRate, effectiveTaxRate)
	}

	expectedTax := float64(11250)
	if expectedTax != tax {
		t.Errorf("Expected Total Tax: %f, Got: %f", expectedTax, tax)
	}
}
