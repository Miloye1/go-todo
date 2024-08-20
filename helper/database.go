package helper

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadDatabase() ([]string, []Record, error) {
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

func WriteToDatabase(recordList []string) error {
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
