package models

import (
	"encoding/json"
	"time"

	"github.com/markphelps/optional"
)

type InvestorResponse struct {
	InterestRate     optional.Float64
	BeginningBalance optional.Float64
	EndingBalance    float64
	YearsHeld        optional.Int
	TaxRate          optional.Float64
	AfterTaxes       bool
	CurrentYear      time.Time
	DesiredAmount    optional.Float64
}

func (res InvestorResponse) String() (string, error) {
	jsonData, err := json.Marshal(res)
	return string(jsonData), err

}
