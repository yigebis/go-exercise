package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(models.Book)
	RemoveBook(int)
	BorrowBook(int, int)
	ReturnBook(int, int)
	ListAvailableBooks()
	ListBorrowedBooks()
}

type Library struct {
	Books    map[int]models.Book
	Members  map[int]models.Member
	Borrowed map[int]int
}

func (L *Library) AddBook(book models.Book) {
	L.Books[book.Id] = book
}

func (L *Library) RemoveBook(bookId int) {
	delete(L.Books, bookId)
}

func (L *Library) BorrowBook(bookId int, memberId int) error {
	book, book_exists := L.Books[bookId]
	_, member_exists := L.Members[memberId]

	if book_exists == false {
		return errors.New("Book not found")
	}

	if member_exists == false {
		return errors.New("Member not found")
	}

	_, is_borrowed := L.Borrowed[bookId]

	if is_borrowed == true {
		return errors.New("Book has been borrowed")
	}

	L.Borrowed[bookId] = memberId
	book.Status = "Borrowed"
	member := L.Members[memberId]
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	fmt.Print("Book borrowed successfully")

	return nil
}

func (L *Library) ReturnBook(bookId, memberId int) error {
	book, book_exists := L.Books[bookId]
	_, member_exists := L.Members[memberId]

	if book_exists == false {
		return errors.New("book not found")
	}

	if member_exists == false {
		return errors.New("Member not found")
	}

	borrower, is_borrowed := L.Borrowed[bookId]

	if is_borrowed == false {
		return errors.New("Book hasn't been borrowed")
	}

	if borrower != memberId {
		return errors.New("Book was not borrowed by this member")
	}

	delete(L.Borrowed, bookId)
	fmt.Println("Book returned succesfully!")
	book.Status = "Available"

	return nil
}

func (L *Library) ListAvailableBooks() []models.Book {
	all_books := L.Books
	var available_books []models.Book

	for id, book := range all_books {
		_, is_borrowed := L.Borrowed[id]
		if !is_borrowed {
			available_books = append(available_books, book)
		}
	}

	return available_books
}

func (L *Library) ListBorrowedBooks(memberId int) []models.Book {
	// all_books := L.Books
	return L.Members[memberId].BorrowedBooks
}
