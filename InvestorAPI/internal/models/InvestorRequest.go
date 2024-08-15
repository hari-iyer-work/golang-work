package models

import (
	"github.com/markphelps/optional"
)

type InvestorRequest struct {
	InterestRate     optional.Float64 `validate:"omitempty,required"`
	BeginningBalance optional.Float64 `validate:"required"`
	YearsHeld        optional.Int     `validate:"required"`
	TaxRate          optional.Float64 `validate:"required,omitempty"`
	AfterTaxes       optional.Bool    `validate:"required"`
	DesiredAmount    optional.Float64 `validate:"required,omitempty"`
}
