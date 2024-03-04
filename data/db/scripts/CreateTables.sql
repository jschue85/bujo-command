DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Journals;
DROP TABLE IF EXISTS DailyLogs;
DROP TABLE IF EXISTS DailyLogItems;

CREATE TABLE Users (id integer PRIMARY KEY, userName string, password string, firstName string, lastName string);
CREATE INDEX IF NOT EXISTS Users_Index on Users(userName);
CREATE TABLE Journals (id integer PRIMARY KEY, year int, userId int, name string, description string, FOREIGN KEY(userId) REFERENCES Users(id) ON DELETE CASCADE);
CREATE INDEX IF NOT EXISTS Journals_Index on Journals (userId);
CREATE INDEX IF NOT EXISTS Journals_Index_year on Journals (userId, year);
CREATE TABLE DailyLogs(id integer PRIMARY KEY, journalId int , month int, day int, dayOfWeek string, FOREIGN KEY (journalId) REFERENCES Journals(id) ON DELETE CASCADE);
CREATE INDEX IF NOT EXISTS DailyLogs_Index on DailyLogs (month, day);

CREATE TABLE DailyLogItems(id integer PRIMARY KEY, logid integer, logType string, details string, status string, FOREIGN KEY (logid) REFERENCES DailyLogs(id) ON DELETE CASCADE);

INSERT INTO Users (userName, password, firstName, lastName) VALUES ("joehschueler@gmail.com", "abc123", "Joe", "Schueler");
INSERT INTO Journals(year, userId, name, description) VALUES (2024, 1, "Home Journal", "Yearly Bullet Journal");
INSERT INTO DailyLogs(journalId, month, day, dayOfWeek) VALUES (1,1,1, "M");

INSERT INTO DailyLogItems (logid, logType, details, status)
VALUES (1,"T", "Organize Toolbox", "New"),
(1, "E", "New Years", "New");

