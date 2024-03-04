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

type JournalModel struct {
	Journal data.Journal
	Daily   DailyModel
}

type DailyDate struct {
	Month int
	Day   int
}

type DailyModel struct {
	Current  data.DailyLog
	Date     time.Time
	Previous DailyDate
	Next     DailyDate
}

func Journal(context echo.Context) error {
	idParam := context.Param("id")
	id, _ := strconv.Atoi(idParam)
	store := &data.BujoStore{}
	model := JournalModel{}
	service := bujo.JournalService{
		Store: store,
	}

	journal, _ := service.GetJournal(id)
	model.Journal = journal
	fmt.Printf("%v\n", journal.Name)
	dailyService := bujo.DailyLogService{
		Store: store,
	}
	today := time.Now()

	daily, _ := dailyService.Get(journal.Id, int(today.Month()), today.Day())
	fmt.Println(daily.Month)
	model.Daily.Current = daily
	model.Daily.Previous = getPrevious(today)

	return context.Render(http.StatusOK, "journal", model)
}

func Daily(context echo.Context) error {
	store := &data.BujoStore{}
	idParam := context.Param("id")
	id, _ := strconv.Atoi(idParam)
	monthParam := context.Param("month")
	month, _ := strconv.Atoi(monthParam)
	dayParam := context.Param("day")
	day, _ := strconv.Atoi(dayParam)

	dailyService := bujo.DailyLogService{
		Store: store,
	}

	daily, _ := dailyService.Get(id, month, day)

	service := bujo.JournalService{
		Store: store,
	}

	journal, _ := service.GetJournal(id)

	date := time.Date(journal.Year, time.Month(daily.Month), daily.Day, 0, 0, 0, 0, time.Local)
	model := DailyModel{
		Current:  daily,
		Previous: getPrevious(date),
		Next:     getNext(date),
	}
	return context.Render(http.StatusOK, "daily", model)
}

func getPrevious(date time.Time) DailyDate {
	previous := date.AddDate(0, 0, -1)
	return DailyDate{
		Month: int(previous.Month()),
		Day:   int(previous.Day()),
	}
}

func getNext(date time.Time) DailyDate {
	next := date.AddDate(0, 0, 1)
	return DailyDate{
		Month: int(next.Month()),
		Day:   int(next.Day()),
	}
}
