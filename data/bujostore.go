package data

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const db_driver = "sqlite3"
const db_name = "bujo.db?_foreign_keys=on"

type BujoStore struct{}

func (s *BujoStore) GetUserById(id int) (User, error) {
	user := User{}
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, userName, firstName, lastName FROM Users WHERE Id=? LIMIT 1", id)
	err = row.Scan(&user.Id, &user.UserName, &user.FirstName, &user.LastName)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}
func (s *BujoStore) GetUserByUserName(userName string) (User, error) {
	user := User{}
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, userName, firstName, lastName FROM Users WHERE userName=? LIMIT 1", userName)
	err = row.Scan(&user.Id, &user.UserName, &user.FirstName, &user.LastName)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}

func (s *BujoStore) AddUser(user User) error {
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		return err
	}
	defer db.Close()

	record, _ := db.Prepare("INSERT INTO User(userName, firstname, lastName) VALUES(?, ?, ?)")
	record.Exec(user.UserName, user.FirstName, user.LastName)

	return nil
}

func (s *BujoStore) GetJournals(userId int) ([]Journal, error) {
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		fmt.Println(err.Error())
		return []Journal{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, userId, name, description, year FROM Journals WHERE userId=?", userId)
	if err != nil {
		fmt.Println(err.Error())
		return []Journal{}, err
	}
	defer rows.Close()

	journals := []Journal{}
	for rows.Next() {
		j := Journal{}
		rows.Scan(&j.Id, &j.UserId, &j.Name, &j.Description, &j.Year)
		journals = append(journals, j)
	}

	return journals, nil
}

func (s *BujoStore) GetJournal(id int) (Journal, error) {
	j := Journal{}
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		fmt.Println(err.Error())
		return j, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, userId, name, description, year FROM Journals WHERE Id=? LIMIT 1", id)
	err = row.Scan(&j.Id, &j.UserId, &j.Name, &j.Description, &j.Year)

	if err != nil {
		fmt.Println(err.Error())
		return j, err
	}

	return j, nil
}

func (s *BujoStore) AddJournal(journal Journal) (int, error) {
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	record, _ := db.Prepare("INSERT INTO Journals(userId, name, description, year) VALUES(?, ?, ?, ?)")
	record.Exec(journal.UserId, journal.Name, journal.Description, journal.Year)
	id := lastInsertedRow(db)
	return id, nil
}
func (s *BujoStore) DeleteJournal(id int) error {
	db, err := sql.Open(db_driver, db_name)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Journals WHERE Id=?", id)

	return err
}

func (s *BujoStore) AddDailyLogs(journal Journal) error {
	db, err := sql.Open("sqlite3", "bujo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var valueStrings []string
	var values []interface{}

	daysInYear := time.Date(journal.Year, 12, 31, 0, 0, 0, 0, time.UTC).YearDay()
	fmt.Println(daysInYear)
	date := time.Date(journal.Year, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < daysInYear; i++ {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		d := date.AddDate(0, 0, i)
		values = append(values, journal.Id, d.Month(), d.Day(), d.Weekday().String())
	}
	sql := fmt.Sprintf("INSERT INTO DailyLogs(journalId, month, day, dayOfWeek) VALUES %s", strings.Join(valueStrings, ", "))

	record, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}

	record.Exec(values...)
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
	daily := DailyLog{
		Month:     month,
		Day:       day,
		JournalId: journalId,
	}
	row := db.QueryRow("SELECT id, dayOfWeek FROM DailyLogs WHERE journalId=? AND month=? AND day =? LIMIT 1", journalId, month, day)
	err = row.Scan(&daily.Id, &daily.DayOfWeek)
	if err != nil {
		fmt.Println(err.Error())
		return DailyLog{}, err
	}

	entries := []DailyEntry{}
	items, err := db.Query("SELECT logType, details, status FROM DailyLogItems WHERE logid=?", daily.Id)
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
	return daily, nil
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

func lastInsertedRow(db *sql.DB) int {
	id := 0
	row := db.QueryRow("SELECT last_insert_rowid()")
	row.Scan(&id)
	return id
}
