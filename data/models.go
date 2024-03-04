package data

type User struct {
	Id        int
	UserName  string
	FirstName string
	LastName  string
}

type Journal struct {
	Id          int
	Year        int
	UserId      int
	Name        string
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
