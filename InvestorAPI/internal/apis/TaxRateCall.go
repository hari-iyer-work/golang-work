package apis

import (
	"encoding/json"
)

type TaxRateCall struct {
	CallToApi
}

func (p TaxRateCall) unMarshalJson(body string) (taxRate any, err error) {
	jsonStr := []byte(body)
	err = json.Unmarshal(jsonStr, taxRate)
	return taxRate, err
}
