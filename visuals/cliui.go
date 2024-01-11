package visuals

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jschue85/bujo-command/bujo"
	"github.com/jschue85/bujo-command/data"
)

const LINE = "================================================================"
const SUB_LINE = "-------------------------------------------------------------"
const CMD_OPEN_DAILY = "d"
const CMD_OPEN_MONTHLY = "m"
const CMD_OPEN_FUTURE = "f"
const CMD_QUIT = "q"

func MainScreen() int {
	journal := data.Journal{
		Id:    1,
		Year:  2024,
		Owner: "Joe Schueler",
	}
	fmt.Printf("%s\n\n", LINE)
	fmt.Printf("Open Daily Log (%s)  Open Monthly Log (%s) Open Future Log (%s) Quit (%s)\n",
		CMD_OPEN_DAILY, CMD_OPEN_MONTHLY, CMD_OPEN_FUTURE, CMD_QUIT)
	fmt.Printf("%s\n\n", LINE)

	var command string
	_, err := fmt.Scanf("%s", &command)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	switch command {
	case CMD_OPEN_DAILY:
		fmt.Printf("Opening Daily ...\n")
		DailyLogScreen(journal)
	case CMD_OPEN_MONTHLY:
		fmt.Printf("Opening Monthly ...\n")
	case CMD_OPEN_FUTURE:
		fmt.Printf("Opening Future ...\n")
	case "q":
		fmt.Printf("Closing ...\n")
		return 0
	}
	return 1
}

func DailyLogScreen(journal data.Journal) {
	dailyService := bujo.DailyLogService{
		Journal: journal,
		Store:   &data.BujoStore{},
	}

	daily, err := dailyService.Get(1, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(LINE)
	fmt.Println()
	fmt.Printf("%v/%v  %s\n", daily.Month, daily.Day, daily.DayOfWeek)
	fmt.Println(SUB_LINE)
	for _, log := range daily.Logs {
		bullet, _ := getBullet(log.Type)
		fmt.Printf("| %s %s \n", bullet, log.Note)
	}
	fmt.Println(SUB_LINE)
	fmt.Printf("Add log (a)  Delete log (d)  back (b)\n")
	fmt.Println(LINE)
	var command string
	_, err = fmt.Scanf("%s", &command)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	switch command {
	case "a":
		fmt.Print("NOTE: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			notePart := strings.SplitN(text, " ", 2)
			noteType, err := getTaskFromBullet(notePart[0])
			if err != nil {
				fmt.Printf("%v\n", err.Error())
				break
			}
			err = dailyService.AddItem(daily.Id, noteType, notePart[1])
			if err != nil {
				fmt.Printf("%v\n", err.Error())
				break
			}
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
		}
		DailyLogScreen(dailyService.Journal)
	case "d":
		fmt.Printf("Deleting ...\n")
	case "b":
		return
	}
}

func getBullet(logType string) (string, error) {
	switch logType {
	case "T":
		return "*", nil
	case "N":
		return "-", nil
	case "E":
		return "O", nil
	case "M":
		return ">", nil
	case "S":
		return "<", nil
	}
	return "", errors.New("invalid log type")
}

func getTaskFromBullet(bullet string) (string, error) {
	switch bullet {
	case "*":
		return "T", nil
	case "_":
		return "N", nil
	case "O":
		return "E", nil
	case ">":
		return "M", nil
	case "<":
		return "S", nil
	}
	return "", errors.New("invalid log type")
}
