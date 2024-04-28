package tax

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var taxLevels = []TaxLevel{
	{"0-150,000", 0.0},
	{"150,001-500,000", 0.0},
	{"500,001-1,000,000", 0.0},
	{"1,000,001-2,000,000", 0.0},
	{"2,000,001 ขึ้นไป", 0.0},
}

var personalDeduction float64 = 60000.0
var kReceiptAmount float64 = 50000.0

func calculateTax(income float64) (float64, []TaxLevel) {
	var tax float64

	switch {
	case income <= 150000:
		tax = 0
	case income >= 150001 && income <= 500000:
		tax = (income - 150000) * 0.1
		taxLevels[1].Tax = tax
	case income >= 500001 && income <= 1000000:
		tax = 35000 + (income-500000)*0.15
		taxLevels[1].Tax = 35000
		taxLevels[2].Tax = (income - 500000) * 0.15
	case income >= 1000001 && income <= 2000000:
		tax = 135000 + (income-1000000)*0.2
		taxLevels[1].Tax = 35000
		taxLevels[2].Tax = 100000
		taxLevels[3].Tax = (income - 1000000) * 0.2
	default:
		tax = 335000 + (income-2000000)*0.35
		taxLevels[1].Tax = 35000
		taxLevels[2].Tax = 100000
		taxLevels[3].Tax = 200000
		taxLevels[4].Tax = (income - 2000000) * 0.35
	}

	return tax, taxLevels
}

func calculateAllowances(person *Person) (float64, float64) {
	var donation, kReceipt float64

	for _, allowance := range person.Allowances {
		switch allowance.AllowanceType {
		case "donation":
			if allowance.Amount > 100000 {
				donation = 100000
			} else {
				donation = allowance.Amount
			}
		case "k-receipt":
			if allowance.Amount > kReceiptAmount {
				kReceipt = kReceiptAmount
			} else {
				kReceipt = allowance.Amount
			}
		}
	}

	return donation, kReceipt
}

func calculateNetIncome(person *Person) float64 {
	donation, kReceipt := calculateAllowances(person)
	netIncome := person.TotalIncome - personalDeduction - donation - kReceipt

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
	tax, taxLevels := calculateTax(netIncome)

	response := Tax{
		Tax:       calculateWht(&person, tax),
		TaxLevels: taxLevels,
	}

	return c.JSON(http.StatusOK, response)
}
