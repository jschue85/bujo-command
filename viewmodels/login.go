package viewmodels

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginForm(context echo.Context) error {
	return context.Render(http.StatusOK, "login", nil)
}

func Login(context echo.Context) error {
	userName := context.FormValue("user")
	WriteCookie(context, "user-name", userName)
	model, _ := getIndexPageData(userName)
	return context.Render(http.StatusOK, "index", model)
}
