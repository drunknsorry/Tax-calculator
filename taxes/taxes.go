package taxes

import (
	"math"

	"github.com/drunknsorry/Tax-calculator/models"
)

// Calculate taxes and return output data, round floats to 2 digits, used for our api
func CalculateTaxes(data *models.TaxBracketResults, year, totalSalary float64) (float64, []float64, []float64, float64) {
	tax := 0.0
	taxesPerBand := make([]float64, len(data.TaxBrackets))
	salaryPerBand := make([]float64, len(data.TaxBrackets))

	// Loop through slice of TaxBrackets
	if totalSalary == 0.0 {
		return 0, taxesPerBand, salaryPerBand, 0.0
	}
	for i, bracket := range data.TaxBrackets {
		if totalSalary <= bracket.Min { // Continue to next loop if salary is lower than min
			continue
		}
		if totalSalary >= bracket.Max && bracket.Max != 0.0 { // If salary is greater than max and max is not zero, do (max - min) * bracket tax rate
			salaryPerBand[i] = (bracket.Max - bracket.Min)
			taxesPerBand[i] = roundFloat((bracket.Max - bracket.Min) * bracket.Rate)
		} else {
			salaryPerBand[i] = (totalSalary - bracket.Min)
			taxesPerBand[i] = roundFloat((totalSalary - bracket.Min) * bracket.Rate) // If salary is less than max, do (salary - min) * bracket tax rate
		}
		tax += taxesPerBand[i]
	}

	effectiveTaxRate := roundFloat(tax / totalSalary)
	tax = roundFloat(tax)
	return tax, taxesPerBand, salaryPerBand, effectiveTaxRate
}

// Function to round to 2 decimal points
func roundFloat(num float64) float64 {
	return math.Round(num*100) / 100
}
