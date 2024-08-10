package handler

import (
	"cms/views/pages"
  "cms/helper"
  "net/http"
	"github.com/labstack/echo/v4"
)

func ShowHomePage(c echo.Context) error {
	return helper.Render(c, http.StatusOK, pages.Home())
}
