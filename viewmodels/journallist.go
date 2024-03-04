package viewmodels

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jschue85/bujo-command/bujo"
	"github.com/jschue85/bujo-command/data"
	"github.com/labstack/echo/v4"
)

func AddJournal(context echo.Context) error {
	userName, _ := ReadCookie(context, "user-name")

	if userName == "" {
		fmt.Println("No User")
		return context.Render(http.StatusOK, "journal-item", nil)
	}
	jName := context.FormValue("name")
	jDesc := context.FormValue("description")
	jYearValue := context.FormValue("year")

	store := &data.BujoStore{}
	model := IndexModel{}
	jYear, _ := strconv.Atoi(jYearValue)

	userService := bujo.UserService{
		Store: store,
	}
	user, _ := userService.GetUser(userName)
	model.CurrentUser = user

	service := bujo.JournalService{
		Store: store,
	}

	id, _ := service.AddJournal(user.Id, jName, jDesc, jYear)
	journal, _ := service.GetJournal(id)
	return context.Render(http.StatusOK, "journal-item", journal)
}

func DeleteJournal(context echo.Context) error {
	idValue := context.Param("id")
	fmt.Printf("DELETING: %v\n", idValue)
	id, _ := strconv.Atoi(idValue)
	store := &data.BujoStore{}
	service := bujo.JournalService{
		Store: store,
	}
	service.DeleteJournal(id)

	return context.NoContent(http.StatusOK)
}
