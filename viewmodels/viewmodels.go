package viewmodels

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jschue85/bujo-command/bujo"
	"github.com/jschue85/bujo-command/data"
	"github.com/labstack/echo/v4"
)

type CurrentOwner struct {
	Owner    string
	Year     int
	Journals []data.Journal
}

func Daily(context echo.Context) error {
	journalValue, _ := readCookie(context, "journal")

	journalService := bujo.JournalService{
		Store: &data.BujoStore{},
	}
	journalId, _ := strconv.Atoi(journalValue)
	j, _ := journalService.GetJournal(journalId)
	dailyService := bujo.DailyLogService{
		Journal: j,
		Store:   &data.BujoStore{},
	}

	daily, err := dailyService.Get(1, 1)
	if err != nil {
		context.Error(err)
	}

	return context.Render(http.StatusOK, "daily", daily)
}

func Index(context echo.Context) error {
	owner, _ := readCookie(context, "owner")
	yearValue, _ := readCookie(context, "year")
	year, _ := strconv.Atoi(yearValue)

	currentOwner := CurrentOwner{
		Owner: owner,
		Year:  year,
	}

	journalService := bujo.JournalService{
		Store: &data.BujoStore{},
	}

	journals, _ := journalService.GetJournals(owner)
	currentOwner.Journals = journals
	writeCookie(context, "journal", fmt.Sprintf("%d", journals[0].Id))
	return context.Render(http.StatusOK, "index", currentOwner)
}

func SelectJournal(context echo.Context) error {
	fmt.Println("Writing cookie")
	writeCookie(context, "owner", "Joe Schueler")
	writeCookie(context, "year", "2024")
	return context.Redirect(http.StatusMovedPermanently, "/")
}

func writeCookie(c echo.Context, name string, value string) error {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return nil
}

func readCookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return cookie.Value, nil
}
