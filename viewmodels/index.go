package viewmodels

import (
	"fmt"
	"net/http"

	"github.com/jschue85/bujo-command/bujo"
	"github.com/jschue85/bujo-command/data"
	"github.com/labstack/echo/v4"
)

type IndexModel struct {
	Title       string
	CurrentUser data.User
	Year        int
	Journals    []data.Journal
}

func Index(context echo.Context) error {
	model := IndexModel{Title: "Bullet Journals"}
	userName, _ := ReadCookie(context, "user-name")

	if userName == "" {
		fmt.Println("No User")
		return context.Render(http.StatusOK, "index", model)
	}
	model, _ = getIndexPageData(userName)

	return context.Render(http.StatusOK, "index", model)
}

func getIndexPageData(userName string) (IndexModel, error) {
	store := &data.BujoStore{}
	model := IndexModel{Title: "Bullet Journals"}

	userService := bujo.UserService{
		Store: store,
	}
	user, _ := userService.GetUser(userName)
	model.CurrentUser = user

	service := bujo.JournalService{
		Store: store,
	}

	journals, _ := service.GetJournals(user.Id)
	model.Journals = journals
	return model, nil
}
