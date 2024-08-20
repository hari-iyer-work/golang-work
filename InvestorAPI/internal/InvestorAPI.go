package investor

import (
	"fmt"
	"time"

	"InvestorAPI/models"
	valid "InvestorAPI/validators"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

func ValidateRequest(req models.InvestorRequest) (isValid bool, err error) {
	validate := validator.New()
	validate.RegisterValidation("CustomValidationFloat", valid.CustomValidationFloat)
	validate.RegisterValidation("CustomValidationInt", valid.CustomValidationInt)
	validate.RegisterValidation("CustomValidationBool", valid.CustomValidationBool)
	err = validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		fmt.Sprintf("Validation error: %s", errors)
		isValid = false
		return
	}
	isValid = true
	return
}

func CompoundInterest(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	var isValid bool
	isValid, err = ValidateRequest(req)
	if err != nil && isValid == false {
		return nil, err
	}
	endingBalance := req.BeginningBalance.MustGet()
	interestRate := (1 + (req.InterestRate.MustGet() / 100))
	taxRate := 1 - (req.TaxRate.MustGet() / 100)
	res = make([]*models.InvestorResponse, req.YearsHeld.MustGet())

	// calculate the compound ending balance after taxes or regardless of taxes
	for years := 0; years < req.YearsHeld.MustGet(); years++ {
		res[years] = &models.InvestorResponse{}
		// copy over similar values from request to response
		err = copier.Copy(res[years], &req)
		if err != nil {
			return nil, err
		}
		if req.AfterTaxes.OrElse(false) == false {
			//non taxable account
			endingBalance = interestRate * endingBalance
			res[years].EndingBalance = endingBalance
			res[years].CurrentYear = time.Now().AddDate(years+1, 0, 0)

		} else if req.AfterTaxes.OrElse(false) == true {
			//taxable account
			taxAmount := ((req.InterestRate.MustGet() / 100) * endingBalance) * taxRate
			afterTaxGains := (endingBalance * interestRate) - taxAmount
			endingBalance = endingBalance + afterTaxGains
			res[years].EndingBalance = endingBalance
			res[years].CurrentYear = time.Now().AddDate(years+1, 0, 0)

		}

	}

	return

}

func GreaterThanDesiredBalance(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	yearsRequired := 0
	var isValid bool
	isValid, err = ValidateRequest(req)
	if err != nil && isValid == false {
		return nil, err
	}
	for req.BeginningBalance.MustGet() < req.DesiredAmount.MustGet() {
		yearsRequired++

		// Set YearsHeld to 1 for each iteration, since we're accumulating year-by-year
		req.YearsHeld.Set(1)

		var yearlyRes []*models.InvestorResponse
		yearlyRes, err = CompoundInterest(req)
		if err != nil {
			return nil, err
		}

		// Update the BeginningBalance for the next iteration to the EndingBalance of this year
		req.BeginningBalance.Set(yearlyRes[0].EndingBalance)

		// Add this year's result to the overall result
		res = append(res, yearlyRes[0])
	}

	return
}
