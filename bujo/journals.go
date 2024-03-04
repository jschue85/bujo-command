package bujo

import (
	"github.com/jschue85/bujo-command/data"
)

type JournalService struct {
	Store data.BujoJournalStore
}

func (j *JournalService) GetJournals(userId int) ([]data.Journal, error) {
	return j.Store.GetJournals(userId)
}

func (j *JournalService) GetJournal(id int) (data.Journal, error) {
	return j.Store.GetJournal(id)
}

func (j *JournalService) AddJournal(userId int, name string, description string, year int) (int, error) {
	journal := data.Journal{
		UserId:      userId,
		Name:        name,
		Description: description,
		Year:        year,
	}
	id, _ := j.Store.AddJournal(journal)
	journal.Id = id

	j.Store.AddDailyLogs(journal)
	return id, nil
}

func (j *JournalService) DeleteJournal(id int) error {
	return j.Store.DeleteJournal(id)
}
