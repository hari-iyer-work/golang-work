package models

import (
	"encoding/json"
	"time"
)

type InvestorResponse struct {
	InterestRate     float64
	BeginningBalance float64
	EndingBalance    float64
	YearsHeld        int
	TaxRate          float64
	AfterTaxes       bool
	CurrentYear      time.Time
	DesiredAmount    float64
}

func (res InvestorResponse) String() (string, error) {
	jsonData, err := json.Marshal(res)
	return string(jsonData), err

}
