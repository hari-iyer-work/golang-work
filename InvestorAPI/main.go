package main

import (
	"fmt"

	investor "github.com/hari-iyer-work/golang-work/tree/main/InvestorAPI/internal"
	models "github.com/hari-iyer-work/golang-work/tree/main/InvestorAPI/internal/models"
	"github.com/markphelps/optional"
)

func main() {
	var emptyValue optional.Int
	var beginngBal optional.Float64
	beginngBal.Set(180000.00)
	var yearsHeld optional.Float64
	yearsHeld.Set(8.9)
	var afterTaxes optional.Bool
	afterTaxes.Set(true)
	var taxRate optional.Float64
	taxRate.Set(10)
	var desiredAmount optional.Float64
	desiredAmount.Set(0)
	request := models.InvestorRequest{InterestRate: yearsHeld, BeginningBalance: beginngBal, YearsHeld: emptyValue, TaxRate: taxRate, AfterTaxes: afterTaxes, DesiredAmount: desiredAmount}
	//requestMillion := models.InvestorRequest{InterestRate: 8, BeginningBalance: 180000.00, YearsHeld: 10, TaxRate: 10, AfterTaxes: true, DesiredAmount: 5000000.00}
	fmt.Println("main app")
	//response, err := investor.CompoundInterest(request)
	response, err := investor.CompoundInterest(request)
	//response, err := investor.GreaterThanDesiredBalance(requestMillion)
	if err != nil {
		fmt.Println(err)
	}
	for _, res := range response {
		if res != nil {
			resactual := *res
			// ... rest of your code

			json, err := resactual.String()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(json)
		}
	}
}
