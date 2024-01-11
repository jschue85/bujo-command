package data

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type BujoStore struct{}

func (s *BujoStore) GetJournals(owner string) ([]Journal, error) {
	j := Journal{
		Id:          1,
		Year:        2024,
		Owner:       owner,
		Description: "Yearly Journal",
	}
	journals := []Journal{j}
	return journals, nil
}

func (s *BujoStore) GetJournal(id int) (Journal, error) {
	j := Journal{
		Id:          id,
		Year:        2024,
		Owner:       "Joe Schueler",
		Description: "Yearly Journal",
	}
	return j, nil
}

func (s *BujoStore) AddJournal(journalId int) error {
	return nil
}
func (s *BujoStore) DeleteJournal(journalId int) error {
	return nil
}

func (s *BujoStore) GetMonthlyLog(month int, year int) {

}

func (s *BujoStore) GetDailyLog(journalId int, month int, day int) (DailyLog, error) {

	db, err := sql.Open("sqlite3", "bujo.db")
	if err != nil {
		fmt.Println(err.Error())
		return DailyLog{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, dayOfWeek FROM DailyLogs WHERE journalId=? AND month=? AND day =? LIMIT 1", journalId, month, day)
	if err != nil {
		fmt.Println(err.Error())
		return DailyLog{}, err
	}
	defer rows.Close()
	daily := DailyLog{
		Month: month,
		Day:   day,
	}
	entries := []DailyEntry{}
	for rows.Next() {
		var id int
		var dayOfWeek string
		rows.Scan(&id, &dayOfWeek)
		daily.DayOfWeek = dayOfWeek
		daily.Id = id
		items, err := db.Query("SELECT logType, details, status FROM DailyLogItems WHERE logid=?", id)
		defer items.Close()
		if err != nil {
			return DailyLog{}, err
		}
		entry := DailyEntry{}
		for items.Next() {
			items.Scan(&entry.Type, &entry.Note, &entry.Status)
			entries = append(entries, entry)
		}
		daily.Logs = entries
	}
	return daily, nil
}

func (s *BujoStore) AddDailyLog(dailyLog DailyLog) error {
	db, err := sql.Open("sqlite3", "bujo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	record, _ := db.Prepare("INSERT INTO DailyLog(journalId, month, day, dayOfWeek) VALUES(?, ?, ?, ?)")

	record.Exec(dailyLog.JournalId, dailyLog.Month, dailyLog.Day, dailyLog.DayOfWeek)
	return nil
}

func (s *BujoStore) AddDailyLogEntry(entry DailyEntry) error {
	db, err := sql.Open("sqlite3", "bujo.db")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	record, _ := db.Prepare("INSERT INTO DailyLogItems (logId, logType, details, status) VALUES(?, ?, ?, ?)")

	record.Exec(entry.DailyId, entry.Type, entry.Note, "New")
	return nil
}

func (s *BujoStore) DeleteDailyLogEntry(entryId int) error {
	db, err := sql.Open("sqlite3", "bujo.db")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM DailyLogItems WHERE id = ?", entryId)

	if err != nil {
		return err
	}
	return nil
}
