package apis

import (
	"encoding/json"
)

type AlphaApiCall struct {
	CallToApi
}

// func (p alphaApiCall)
func (p AlphaApiCall) unmarshalJson(body []byte) (interestRate interface{}, err error) {
	err = json.Unmarshal(body, interestRate)
	return interestRate, err
}
