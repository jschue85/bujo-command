package bujo

import (
	"github.com/jschue85/bujo-command/data"
)

type DailyLogService struct {
	Journal data.Journal
	Store   data.BujoDailyLogStore
}

func (s DailyLogService) Get(month int, day int) (data.DailyLog, error) {
	daily, err := s.Store.GetDailyLog(s.Journal.Id, month, day)
	if err != nil {
		return daily, err
	}

	return daily, nil
}

func (s DailyLogService) Add(month int, day int) error {
	log := data.DailyLog{
		JournalId: s.Journal.Id,
		Month:     month,
		Day:       day,
	}
	return s.Store.AddDailyLog(log)
}

func (s DailyLogService) AddItem(id int, logType string, details string) error {

	entry := data.DailyEntry{
		DailyId: id,
		Type:    logType,
		Note:    details,
	}
	return s.Store.AddDailyLogEntry(entry)
}
