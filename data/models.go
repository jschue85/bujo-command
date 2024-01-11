package data

type Journal struct {
	Id          int
	Year        int
	Owner       string
	Description string
}

type DailyLog struct {
	Id        int
	JournalId int
	Month     int
	Day       int
	DayOfWeek string
	Logs      []DailyEntry
}

type DailyEntry struct {
	DailyId int
	Type    string
	Note    string
	Status  string
}
