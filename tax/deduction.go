package tax

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateDeductionHandler(c echo.Context) error {
	var request DeductionLimits

	if err := c.Bind(&request); err != nil {
		return err
	}

	personalDeduction = request.Amount

	return c.JSON(http.StatusOK, map[string]float64{
		"personalDeduction": personalDeduction,
	})
}
func UpdateKReceiptHandler(c echo.Context) error {
    var request DeductionLimits
    
	if err := c.Bind(&request); err != nil {
        return err
    }

    kReceiptAmount = request.Amount

    return c.JSON(http.StatusOK, map[string]float64{
        "kReceipt": kReceiptAmount,
    })
}
