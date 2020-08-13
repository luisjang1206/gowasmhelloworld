package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e := echo.New()
	e.Use(middleware.Static("./"))
	e.File("/", "./index.html")
	e.Start(":8000")
}
