DROP TABLE IF EXISTS Journals;
DROP TABLE IF EXISTS DailyLogs;
DROP TABLE IF EXISTS DailyLogItems;

CREATE TABLE Journals (id integer PRIMARY KEY, year int, owner string, description string);
CREATE INDEX IF NOT EXISTS Journals_Index on Journals (year);
CREATE TABLE DailyLogs(id integer PRIMARY KEY, journalId int , month int, day int, dayOfWeek string, FOREIGN KEY (journalId) REFERENCES Journals(id));
CREATE INDEX IF NOT EXISTS DailyLogs_Index on DailyLogs (month, day);

CREATE TABLE DailyLogItems(id integer PRIMARY KEY, logid integer, logType string, details string, status string, FOREIGN KEY (logid) REFERENCES DailyLogs(id));


INSERT INTO Journals(year, owner, description) VALUES (2024, "Joe Schueler", "Yearly Bullet Journal");
INSERT INTO DailyLogs(journalId, month, day, dayOfWeek) VALUES (1,1,1, "M");

INSERT INTO DailyLogItems (logid, logType, details, status)
VALUES (1,"T", "Organize Toolbox", "New"),
(1, "E", "New Years", "New");

