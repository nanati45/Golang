package models

type  Member struct {
	MemberID int
	Name string 
	BorrowedBooks []Book
}