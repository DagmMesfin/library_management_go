package services

import "library-management/models"

type LibraryManager interface {
	AddMember(member models.Member)
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListMembers() map[int]models.Member
	ListBorrowedBooks(memberID int) []models.Book
}
