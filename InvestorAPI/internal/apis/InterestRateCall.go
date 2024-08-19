package apis

import (
	"encoding/json"
)

type InterestRateCall struct {
	CallToApi
}

func (p InterestRateCall) unMarshalJson(body json) (interestRate any, err error) {
	err = json.Unmarshal(body, interestRate)
	return interestRate, err
}
