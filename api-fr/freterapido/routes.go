package freterapido

import "github.com/labstack/echo"

func Routes(e *echo.Echo) {

	e.POST("/quote", handleQuote)
	e.GET("/metrics", GetMetrics)

	e.Logger.Fatal(e.Start(":9090"))

}
