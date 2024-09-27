package freterapido

import "github.com/labstack/echo"

func Routes(e *echo.Echo) {

	e.POST("/quote", handleQuote)
	e.GET("/metrics", GetMetrics)
	e.GET("/metrics/carrierprice", GetMetricsByCarrier)
	e.GET("metrics/cheaper", GetCheaperQuote)
	e.GET("/metrics/expensive", GetExpensiveQuote)

	e.Logger.Fatal(e.Start(":9090"))

}
