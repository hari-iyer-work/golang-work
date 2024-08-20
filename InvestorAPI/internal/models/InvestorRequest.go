package models

import (
	"github.com/markphelps/optional"
)

type InvestorRequest struct {
	InterestRate     optional.Float64 `validate:"CustomValidationFloat"`
	BeginningBalance optional.Float64 `validate:"CustomValidationFloat"`
	YearsHeld        optional.Int     `validate:"CustomValidationInt"`
	TaxRate          optional.Float64 `validate:"CustomValidationFloat"`
	AfterTaxes       optional.Bool    `validate:"CustomValidationBool"`
	DesiredAmount    optional.Float64 `validate:"CustomValidationFloat"`
}
