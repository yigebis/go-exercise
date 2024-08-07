package models

//define member
type Member struct {
	Id            int
	Name          string
	BorrowedBooks []Book
}
