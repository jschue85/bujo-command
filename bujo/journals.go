package bujo

import (
	"github.com/jschue85/bujo-command/data"
)

type JournalService struct {
	Store data.BujoJournalStore
}

func (j *JournalService) GetJournals(owner string) ([]data.Journal, error) {
	return j.Store.GetJournals(owner)
}

func (j *JournalService) GetJournal(id int) (data.Journal, error) {
	return j.Store.GetJournal(id)
}
