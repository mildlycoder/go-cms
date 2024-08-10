package main

import (
	"cms/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/public", "public")
	e.GET("/", handler.ShowHomePage)
	e.GET("/login", handler.ShowloginPage)
	e.GET("/signup", handler.ShowSignupPage)

	e.Logger.Fatal(e.Start(":3000"))
}
