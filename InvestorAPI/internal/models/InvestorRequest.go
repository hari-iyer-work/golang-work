package models

type InvestorRequest struct {
	InterestRate     float64
	BeginningBalance float64
	YearsHeld        int
	TaxRate          float64
	AfterTaxes       bool
	DesiredAmount    float64
}

