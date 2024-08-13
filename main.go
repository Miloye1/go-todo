package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	id   int
	task string
	done bool
}

func (r Record) toString() []string {
	return []string{fmt.Sprintf("%v", r.id), r.task, fmt.Sprintf("%v", r.done)}
}

func main() {
	for {
		fmt.Println("")
		fmt.Println("1. Show all todos")
		fmt.Println("2. Add new todo")
		fmt.Println("3. Mark todo as done")
		fmt.Println("")
		fmt.Println("q. Quit")

		fmt.Print("Select option: ")
		userInput, scanErr := scanUserInput()

		if scanErr != nil {
			fmt.Printf("Error while reading user input, %v", scanErr)
			break
		}

		switch strings.ToLower(userInput) {
		case "q":
			fmt.Println("User exited")
			os.Exit(0)
		case "1":
			showTodos()
		case "2":
			addNewTodoToDatabase()
		case "3":
			fmt.Println("User chose 3")
		default:
			fmt.Println(userInput)
		}
	}
}

func showTodos() {
	header, parsedRecords, err := readDatabase()

	if err != nil {
		return
	}

	prettyPrint(header)

	for _, record := range parsedRecords {
		prettyPrint(record.toString())
	}
}

func addNewTodoToDatabase() {
	fmt.Println("Add new todo: ")
	task, scanErr := scanUserInput()

	if scanErr != nil {
		return
	}

	record := Record{
		id:   1,
		task: task,
		done: false,
	}

	_, parsedRecords, err := readDatabase()

	if os.IsNotExist(err) {
		fmt.Println("Creating new database...")

		header := []string{"Id", "Todo", "Done"}

		if err := writeToDatabase(header); err != nil {
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

	if err := writeToDatabase(recordList); err != nil {
		return
	}

	fmt.Println("Added new todo to the database")
}

func readDatabase() ([]string, []Record, error) {
	file, err := os.Open("./.data/db.csv")

	if err != nil {
		fmt.Println("Todo database doesn't exist")
		return nil, nil, err
	}

	defer func() {
		err := file.Close()

		if err != nil {
			fmt.Printf("Error while closing the file, %v\n", err)
			return
		}
	}()

	reader := csv.NewReader(file)
	records, readerErr := reader.ReadAll()

	if readerErr != nil {
		fmt.Printf("Error while opening the database, %v\n", readerErr)
		return nil, nil, readerErr
	}

	header := records[0]
	parsedRecords, parseErr := parseRecords(records[1:])

	if parseErr != nil {
		fmt.Printf("Error while parsing the records, %v\n", parseErr)
		return nil, nil, parseErr
	}

	return header, parsedRecords, nil
}

func writeToDatabase(recordList []string) error {
	file, err := os.OpenFile("./.data/db.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Todo database doesn't exist")
		return err
	}

	defer func() {
		err := file.Close()

		if err != nil {
			fmt.Printf("Error while closing the file, %v\n", err)
			return
		}
	}()

	writer := csv.NewWriter(file)

	writer.Write(recordList)
	writer.Flush()

	writeErr := writer.Error()

	if writeErr != nil {
		fmt.Printf("Error while writing to the file, %v\n", writeErr)
		return writeErr
	}

	return nil
}

func scanUserInput() (string, error) {
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
