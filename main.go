package main

import (
	"cms/views/pages"
	"html/template"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func Render( c echo.Context, statusCode int, t templ.Component) error {

  c.Response().Writer.WriteHeader(statusCode)
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
  return t.Render(c.Request().Context(), c.Response().Writer)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("./views/**/*.html")),
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, pages.Home())
	})
	
	e.Logger.Fatal(e.Start(":3000"))
}
