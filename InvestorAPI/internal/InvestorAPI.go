package investor

import (
	"time"

	"github.com/hari-iyer-work/golang-work/tree/main/InvestorAPI/internal/models"
	"github.com/jinzhu/copier"
)

func CompoundInterest(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	endingBalance := req.BeginningBalance
	interestRate := (1 + (req.InterestRate / 100))
	taxRate := 1 - (req.TaxRate / 100)
	res = make([]*models.InvestorResponse, req.YearsHeld)

	// calculate the compound ending balance after taxes or regardless of taxes
	for years := 0; years < req.YearsHeld; years++ {
		res[years] = &models.InvestorResponse{}
		// copy over similar values from request to response
		err = copier.Copy(res[years], &req)
		if err != nil {
			return nil, err
		}
		if req.AfterTaxes == false {
			//non taxable account
			endingBalance = interestRate * endingBalance
			res[years].EndingBalance = endingBalance
			res[years].CurrentYear = time.Now().AddDate(years+1, 0, 0)

		} else if req.AfterTaxes == true {
			//taxable account
			taxAmount := ((req.InterestRate / 100) * endingBalance) * taxRate
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

	for req.BeginningBalance < req.DesiredAmount {
		yearsRequired++

		// Set YearsHeld to 1 for each iteration, since we're accumulating year-by-year
		req.YearsHeld = 1

		var yearlyRes []*models.InvestorResponse
		yearlyRes, err = CompoundInterest(req)
		if err != nil {
			return nil, err
		}

		// Update the BeginningBalance for the next iteration to the EndingBalance of this year
		req.BeginningBalance = yearlyRes[0].EndingBalance

		// Add this year's result to the overall result
		res = append(res, yearlyRes[0])
	}

	return
}
