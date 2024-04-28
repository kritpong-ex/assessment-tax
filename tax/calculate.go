package tax

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var personalDeduction float64 = 60000.0

func calculateTax(income float64) float64 {
	var tax float64

	switch {
	case income <= 150000:
		tax = 0
	case income >= 150001 && income <= 500000:
		tax = (income - 150000) * 0.1
	case income >= 500001 && income <= 1000000:
		tax = 35000 + (income-500000)*0.15
	case income >= 1000001 && income <= 2000000:
		tax = 135000 + (income-1000000)*0.2
	default:
		tax = 335000 + (income-2000000)*0.35

	}

	return tax
}

func calculateAllowances(person *Person) float64 {
	var donation float64

	for _, allowance := range person.Allowances {
		switch allowance.AllowanceType {
		case "donation":
			if allowance.Amount > 100000 {
				donation = 100000
			} else {
				donation = allowance.Amount
			}
		}
	}

	return donation
}

func calculateNetIncome(person *Person) float64 {
	donation := calculateAllowances(person)
	netIncome := person.TotalIncome - personalDeduction - donation

	if netIncome < 0 {
        return 0
    }

	return netIncome
}

func calculateWht(person *Person, tax float64) float64 {
	if person.WHT > 0 {
		tax -= person.WHT
	}
	
	if tax < 0 {
		tax = 0
	}

	return tax
}

func CalculateTaxHandler(c echo.Context) error {
	var person Person

	if err := c.Bind(&person); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&person); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	netIncome := calculateNetIncome(&person)
	tax := calculateTax(netIncome)

	response := Tax{
		Tax:       calculateWht(&person, tax),
	}

	return c.JSON(http.StatusOK, response)
}
