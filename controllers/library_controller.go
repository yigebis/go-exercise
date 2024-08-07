package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

var library *services.Library

func init() {
	library = &services.Library{
		Books:    make(map[int]models.Book),
		Members:  make(map[int]models.Member),
		Borrowed: make(map[int]int),
	}

	sample_members := []models.Member{
		{Id: 0, Name: "Alemu", BorrowedBooks: []models.Book{}},
		{Id: 1, Name: "Alice", BorrowedBooks: []models.Book{}},
		{Id: 2, Name: "Bob", BorrowedBooks: []models.Book{}},
		{Id: 3, Name: "Charlie", BorrowedBooks: []models.Book{}},
		{Id: 4, Name: "Diana", BorrowedBooks: []models.Book{}},
		{Id: 5, Name: "Eve", BorrowedBooks: []models.Book{}},
	}

	for i, member := range sample_members {
		library.Members[i] = member
	}

	fmt.Println(library.Members)
}

func Run() {
	for {
		fmt.Println("Choose an option (Choose 0 for exit)\n")

		var choice int
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List books borrowed by a member")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			book := QueryBook()
			library.AddBook(book)

		case 2:
			bookId := QueryBookId()
			library.RemoveBook(bookId)

		case 3:
			bookId := QueryBookId()
			memberId := QueryMemberId()

			library.BorrowBook(bookId, memberId)

		case 4:
			bookId := QueryBookId()
			memberId := QueryMemberId()

			library.ReturnBook(bookId, memberId)

		case 5:
			available_books := library.ListAvailableBooks()
			fmt.Println("ID...Title...Author")

			for _, book := range available_books {
				fmt.Printf("%d...%s...%s\n", book.Id, book.Title, book.Author)
			}

		case 6:
			memberId := QueryMemberId()

			borrowed_books := library.ListBorrowedBooks(memberId)
			fmt.Println("ID...Title...Author")

			for _, book := range borrowed_books {
				fmt.Printf("%d...%s...%s\n", book.Id, book.Title, book.Author)
			}
		case 0:
			return

		default:
			fmt.Println("Incorrect Input! try again")
		}
	}
}

func QueryBookId() int {
	var bookId int
	fmt.Println("Enter Book ID")
	fmt.Scan(&bookId)

	return bookId
}

func QueryMemberId() int {
	var memberId int
	fmt.Println("Enter Member ID")
	fmt.Scan(&memberId)

	return memberId
}

func QueryBook() models.Book {
	var bookId int
	var title, author string

	fmt.Println("Enter the title of the book")
	fmt.Scan(&title)

	fmt.Println("Enter the author of the book")
	fmt.Scan(&author)

	fmt.Println("Assign the book an Id")
	fmt.Scan(&bookId)

	book := models.Book{Id: bookId, Title: title, Author: author, Status: "Available"}

	return book

}
