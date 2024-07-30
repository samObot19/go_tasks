package main

import (
	"fmt"
	"library_management/services"
	"library_management/controllers"
)

func main() {
	manager := services.NewLibraryManager()

	for {
		controllers.PrintMenu()
		choice := controllers.ReadUserInput()

		switch choice {
		case "1":
			controllers.AddBook(manager)
		case "2":
			controllers.AddMember(manager)
		case "3":
			controllers.BorrowBook(manager)
		case "4":
			controllers.ReturnBook(manager)
		case "5":
			controllers.ListAvailableBooks(manager)
		case "6":
			controllers.ListBorrowedBooks(manager)
		case "7":
			fmt.Println("good buy!!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
