package data

type BujoUserStore interface {
	GetUserById(id int) (User, error)
	GetUserByUserName(userName string) (User, error)
	AddUser(user User) error
}
type BujoJournalStore interface {
	GetJournal(journalId int) (Journal, error)
	GetJournals(userId int) ([]Journal, error)
	AddJournal(journal Journal) (int, error)
	DeleteJournal(journalId int) error
	AddDailyLogs(journal Journal) error
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
	AddDailyLogEntry(entry DailyEntry) error
	DeleteDailyLogEntry(entryId int) error
}
