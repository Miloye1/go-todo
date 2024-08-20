package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowTodos() {
	header, parsedRecords, err := ReadDatabase()

	if err != nil {
		return
	}

	prettyPrint(header)

	for _, record := range parsedRecords {
		prettyPrint(record.toString())
	}
}

func AddNewTodoToDatabase() {
	fmt.Println("Add new todo: ")
	task, scanErr := ScanUserInput()

	if scanErr != nil {
		return
	}

	record := Record{
		id:   1,
		task: task,
		done: false,
	}

	_, parsedRecords, err := ReadDatabase()

	if os.IsNotExist(err) {
		fmt.Println("Creating new database...")

		header := []string{"Id", "Todo", "Done"}

		if err := WriteToDatabase(header); err != nil {
			return
		}
	}

	if len(parsedRecords) > 0 {
		var maxId int = -1 << 63

		for _, record := range parsedRecords {
			if record.id > maxId {
				maxId = record.id
			}
		}

		record.id = maxId + 1
	}

	recordList := record.toString()

	if err := WriteToDatabase(recordList); err != nil {
		return
	}

	fmt.Println("Added new todo to the database")
}

func ScanUserInput() (string, error) {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while getting user input %v\n", err)
		return "", err
	}

	return input, nil
}

func parseRecords(records [][]string) ([]Record, error) {
	parsedRecords := []Record{}

	for _, row := range records {
		id, idErr := strconv.Atoi(row[0])

		if idErr != nil {
			fmt.Printf("Error converting string to int, %v", idErr)
			return nil, idErr
		}

		done, doneErr := strconv.ParseBool(row[2])

		if doneErr != nil {
			fmt.Printf("Error converting string to bool, %v", doneErr)
			return nil, doneErr
		}

		newRecord := Record{
			id:   id,
			task: row[1],
			done: done,
		}

		parsedRecords = append(parsedRecords, newRecord)
	}

	return parsedRecords, nil
}

func prettyPrint(records []string) {
	rowToPrint := strings.Join(records, " | ")
	fmt.Println(rowToPrint)
}
