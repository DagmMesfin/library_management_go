package controllers

import (
	"bufio"
	"fmt"
	"library-management/models"
	"library-management/services"
	"os"
	"strconv"
	"strings"
	"time"
)

type LibraryController struct {
	services services.Library
	reader   *bufio.Reader
}

func NewLibraryController(service services.Library) *LibraryController {
	return &LibraryController{
		services: service,
		reader:   bufio.NewReader(os.Stdin),
	}
}

func (libcon *LibraryController) AddMember(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Print("\tEnter Member ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid Member ID")
		fmt.Println("===================================")
		return
	}

	fmt.Print("\tEnter Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	Member := models.Member{
		ID:            id,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}

	libcon.services.AddMember(Member)
	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second)
	fmt.Println("  Success: Member added successfully.")
	fmt.Println("===================================")

}

func (libcon *LibraryController) AddBook(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("======================================")
	fmt.Print("\tEnter book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		fmt.Println("===================================")
		return
	}

	fmt.Print("\tEnter title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("\tEnter author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	libcon.services.AddBook(book)
	fmt.Println("\nWaiting...")
	time.Sleep(3 * time.Second)
	fmt.Println("Book added successfully.")
	fmt.Println("===================================")

}

func (c *LibraryController) RemoveBook(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("======================================")
	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		fmt.Println("======================================")
		return
	}

	c.services.RemoveBook(id)

	fmt.Println("Book removed successfully.")
}

func (c *LibraryController) ListAvailableBooks() {
	fmt.Println()
	fmt.Println("======================================")
	books := c.services.ListAvailableBooks()
	fmt.Println("Available Books:")
	fmt.Println("\tID\tTitle\tAuthor")
	for _, book := range books {
		fmt.Printf("\t%d\t%s\t%s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) ListBorrowedBooks(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("======================================")
	fmt.Print("Enter member ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid member ID")
		fmt.Println("======================================")
		return
	}

	books, err := c.services.ListBorrowedBooks(id)

	if err != nil {
		if err.Error() == "member not found" {
			fmt.Println("The Member is not found.")
		} else if err.Error() == "no borrowed books" {
			fmt.Println("No borrowed books found for this member.")
		}
		fmt.Println("======================================")
		return
	}

	fmt.Println("Borrowed Books:")
	fmt.Println("\tID\tTitle\tAuthor")

	for _, book := range books {
		fmt.Printf("\t%d\t%s\t%s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) BorrowBook(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("======================================")
	fmt.Print("Enter book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(bookIDStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		fmt.Println("======================================")
		return
	}

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(memberIDStr))
	if err != nil {
		fmt.Println("Invalid member ID")
		fmt.Println("======================================")
		return
	}

	err = c.services.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
	fmt.Println("======================================")
}

func (c *LibraryController) ReturnBook(reader *bufio.Reader) {
	fmt.Println()
	fmt.Println("======================================")
	fmt.Print("Enter book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(bookIDStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		fmt.Println("======================================")
		return
	}

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(memberIDStr))
	if err != nil {
		fmt.Println("Invalid member ID")
		fmt.Println("======================================")
		return
	}

	err = c.services.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
	fmt.Println("======================================")
}
