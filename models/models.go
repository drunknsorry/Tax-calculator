package models

// Struct defining the data structure for the json payload for the api consumer
type TaxBracketResults struct {
	TaxBrackets []Brackets `json:"tax_brackets"`
}

// Struct defining the slice inside of TaxBrackets from TaxBracketResults for the api consumer
type Brackets struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Rate float64 `json:"rate"`
}

// Struct defining the json response structure for our api
type TaxResponse struct {
	Salary           float64   `json:"Salary"`
	TotalTaxesOwed   float64   `json:"Total Taxes Owed"`
	TaxesPerBand     []float64 `json:"Taxes Per Band"`
	SalaryPerTaxBand []float64 `json:"Salary Per Tax Band"`
	EffectiveTaxRate float64   `json:"Effective Tax Rate"`
}
