package bujo

import (
	"github.com/jschue85/bujo-command/data"
)

type DailyLogService struct {
	Store data.BujoDailyLogStore
}

func (s DailyLogService) Get(journalId int, month int, day int) (data.DailyLog, error) {
	daily, err := s.Store.GetDailyLog(journalId, month, day)
	if err != nil {
		return daily, err
	}

	return daily, nil
}

func (s DailyLogService) AddItem(id int, logType string, details string) error {

	entry := data.DailyEntry{
		DailyId: id,
		Type:    logType,
		Note:    details,
	}
	return s.Store.AddDailyLogEntry(entry)
}
