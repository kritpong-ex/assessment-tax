package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kritpong-ex/assessment-tax/tax"
)

func main() {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.POST("/tax/calculations", tax.CalculateTaxHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
