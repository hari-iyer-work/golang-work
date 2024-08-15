package models

type InvestorRequest struct {
	InterestRate     *float64 `validate:"required"`
	BeginningBalance *float64 `validate:"required"`
	YearsHeld        *int     `validate:"required"`
	TaxRate          float64
	AfterTaxes       *bool `validate:"required"`
	DesiredAmount    float64
}
