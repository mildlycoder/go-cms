package handler

import (
	"cms/helper"
	"cms/models"
	"cms/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)


type User struct {
    Email         string `json:"email"`
    Name         string `json:"name"`
}

func ShowSignupPage(c echo.Context) error {
	return helper.Render(c, http.StatusOK, pages.Signup())
}
func BeginAuth(c echo.Context) error {
	provider := c.Param("provider")
  if provider == "" {
   return c.String(http.StatusBadRequest, "Provider not specified")
  }

  q := c.Request().URL.Query()
  q.Add("provider", c.Param("provider"))
  c.Request().URL.RawQuery = q.Encode()

  req := c.Request()
  res := c.Response().Writer
  if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
   return c.JSON(http.StatusOK, gothUser)
  }
  gothic.BeginAuthHandler(res, req)
  return nil
}

func CallbackAuth(c echo.Context) error {
    req := c.Request()
    res := c.Response().Writer
    user, err := gothic.CompleteUserAuth(res, req)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }
    
    // Store user information in session
    session, _ := gothic.Store.Get(req, "user-session")
    session.Values["email"] = user.Email
    session.Values["name"] = user.Name
    session.Save(req, res)

    // Redirect to profile page
    return c.Redirect(http.StatusSeeOther, "/profile")
}

func ShowProfilePage(c echo.Context) error {
    req := c.Request()
    session, _ := gothic.Store.Get(req, "user-session")
    
    email, ok := session.Values["email"].(string)
    if !ok {
        return c.Redirect(http.StatusSeeOther, "/login")
    }
    
    name, ok := session.Values["name"].(string)
    if !ok {
        return c.Redirect(http.StatusSeeOther, "/login")
    }

    user := models.User{
        Email: email,
        Name: name,
    }

    return helper.Render(c, http.StatusOK, pages.Profile(user))
}
