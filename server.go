package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "QR Code creation service. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

func getQRHandler(c echo.Context) error {
	qr := getQR()
	return c.File(qr)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.GET("/qr", getQRHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
