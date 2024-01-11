package data

type BujoJournalStore interface {
	GetJournal(id int) (Journal, error)
	GetJournals(owner string) ([]Journal, error)
	AddJournal(journalId int) error
	DeleteJournal(journalId int) error
}

type BujoFutureLogStore interface {
	GetFuture(journalId int)
	AddFutureLog()
}

type BujoMonthlyStore interface {
	GetMonthly(journalId int, month int)
	AddMonthlyLog(journalId int, month int, day int)
}

type BujoDailyLogStore interface {
	GetDailyLog(journalId int, month int, day int) (DailyLog, error)
	AddDailyLog(log DailyLog) error
	AddDailyLogEntry(entry DailyEntry) error
	DeleteDailyLogEntry(entryId int) error
}
