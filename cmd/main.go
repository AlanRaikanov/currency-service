package main

import (
	"./handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/currency", controllers.SaveCurrencyHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
