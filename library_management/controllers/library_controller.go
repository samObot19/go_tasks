package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

// PrintMenu displays the menu options for the library management system
func PrintMenu() {
	fmt.Println("\nLibrary Management System")
	fmt.Println("1. Add a book")
	fmt.Println("2. Add a member")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List available books")
	fmt.Println("6. List borrowed books")
	fmt.Println("7. Exit")
	fmt.Print("Enter your choice: ")
}

// ReadUserInput reads user input from the console and returns it as a string
func ReadUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

// ReadInt reads an integer value from the console and returns it
func ReadInt() int {
	var value int
	fmt.Scan(&value)
	return value
}

// ReadString reads a string value from the console and returns it
func ReadString() string {
	var input string
	fmt.Scan(&input)
	return input
}

// AddBook adds a new book to the library management system
func AddBook(manager services.LibraryManager) {
	fmt.Print("Enter book ID: ")
	id := ReadInt()
	fmt.Print("Enter book title: ")
	title := ReadString()
	fmt.Print("Enter book author: ")
	author := ReadString()
	book := &models.Book{ID: id, Title: title, Author: author, Status: "available"}
	manager.AddBook(book)
	fmt.Println("Book added successfully.")
}

// AddMember adds a new member to the library management system
func AddMember(manager services.LibraryManager) {
	fmt.Print("Enter member ID: ")
	id := ReadInt()
	fmt.Print("Enter member name: ")
	name := ReadString()
	member := &models.Member{ID: id, Name: name}
	manager.AddMember(member)
	fmt.Println("Member added successfully.")
}

func BorrowBook(manager services.LibraryManager){
	fmt.Println("Enter member ID: ")
	memberid := ReadInt()
	fmt.Println("Enter book ID: ")
	bookid := ReadInt()
	err := manager.BorrowBook(memberid, bookid)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("You borrowed the book with id ", bookid, " succesfully!")

}

func ReturnBook(manager services.LibraryManager){
	fmt.Println("Enter book ID: ")
	bookid := ReadInt()
	fmt.Println("Enter your member ID: ")
	memberid := ReadInt()
	err := manager.ReturnBook(bookid, memberid)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("you returned back successfully!")
}

func ListAvailableBooks(manager services.LibraryManager) {
	books := manager.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}

	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func ListBorrowedBooks(manager services.LibraryManager) {
	fmt.Println("Enter the member id")
	memberid := ReadInt()
	borrowedBooks := manager.ListBorrowedBooks(memberid)
	if len(borrowedBooks) == 0 {
		fmt.Println("No books are currently borrowed.")
		return
	}

	fmt.Println("\nBorrowed Books:")
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}