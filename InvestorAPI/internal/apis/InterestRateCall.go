package apis

import (
	"encoding/json"
)

type InterestRateCall struct {
	CallToApi
}

func (p InterestRateCall) UnmarshalJson(body []byte) (interestRate interface{}, err error) {
	err = json.Unmarshal(body, interestRate)
	return interestRate, err
}
