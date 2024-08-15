package main

import (
	"fmt"

	investor "github.com/hari-iyer-work/golang-work/tree/main/InvestorAPI/internal"
	models "github.com/hari-iyer-work/golang-work/tree/main/InvestorAPI/internal/models"
)

func main() {
	var emptyValue *int = nil
	var beginngBal float64 = 180000.00
	yearsHeld := float64(8.9)
	afterTaxes := true
	request := models.InvestorRequest{InterestRate: &(yearsHeld), BeginningBalance: &beginngBal, YearsHeld: emptyValue, TaxRate: 10, AfterTaxes: &afterTaxes, DesiredAmount: 0}
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
