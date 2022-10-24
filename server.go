package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func initServer() *echo.Echo {
	e := echo.New()

	// middle ware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error { return HomePage(c, DB) })

	return e
}

func HomePage(c echo.Context, db *gorm.DB) error {
	return c.String(http.StatusOK, "<h1>Home Page</h1>")
}
