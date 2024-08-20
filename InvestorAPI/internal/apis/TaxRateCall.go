package apis

import (
	"encoding/json"
)

type TaxRateCall struct {
	CallToApi
}

func (p TaxRateCall) UnmarshalJson(body []byte) (taxRate interface{}, err error) {
	err = json.Unmarshal(body, taxRate)
	return taxRate, err
}
