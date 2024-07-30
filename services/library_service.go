package services

import (
	"errors"
	"library-management/models"
)

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewInstanceLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (lib *Library) AddMember(member models.Member) {
	lib.members[member.ID] = member
}

func (lib *Library) AddBook(book models.Book) {
	lib.books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.books, bookID)
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, book_exist := lib.books[bookID]

	if !book_exist {
		return errors.New("book not found")
	} else if book.Status == "Borrowed" {
		return errors.New("the book is borrowed")
	}

	member, member_exist := lib.members[memberID]

	if !member_exist {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	lib.books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.members[memberID] = member

	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, book_exist := lib.books[bookID]

	if !book_exist {
		return errors.New("book not found")
	} else if book.Status == "Available" {
		return errors.New("the book wasn't borrowed")
	}

	member, member_exist := lib.members[memberID]

	if !member_exist {
		return errors.New("member not found")
	}

	book.Status = "Available"
	lib.books[bookID] = book

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	lib.members[memberID] = member

	return nil

}

func (lib *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book

	for _, book := range lib.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, member_exist := lib.members[memberID]

	if !member_exist {
		return []models.Book{}, errors.New("member not found")
	}

	if len(member.BorrowedBooks) == 0 {
		return []models.Book{}, errors.New("no borrowed books")
	}

	return member.BorrowedBooks, nil
}
