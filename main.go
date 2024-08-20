package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Miloye1/go-todo/helper"
)

func main() {
	for {
		fmt.Println("")
		fmt.Println("1. Show all todos")
		fmt.Println("2. Add new todo")
		fmt.Println("3. Mark todo as done")
		fmt.Println("")
		fmt.Println("q. Quit")

		fmt.Print("Select option: ")
		userInput, scanErr := helper.ScanUserInput()

		if scanErr != nil {
			fmt.Printf("Error while reading user input, %v", scanErr)
			break
		}

		switch strings.ToLower(userInput) {
		case "q":
			fmt.Println("User exited")
			os.Exit(0)
		case "1":
			helper.ShowTodos()
		case "2":
			helper.AddNewTodoToDatabase()
		case "3":
			fmt.Println("User chose 3")
		default:
			fmt.Println(userInput)
		}
	}
}
