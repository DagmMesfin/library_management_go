package main

import (
	"bufio"
	"fmt"
	"library-management/controllers"
	"library-management/services"
	"os"
	"strings"
	"time"
)

func main() {
	library := services.NewInstanceLibrary()
	c := controllers.NewLibraryController(*library)

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("\033[H\033[2J") //clear the console
		fmt.Println("📚 Welcome to the Library Management System 📚")
		fmt.Println("==============================")
		fmt.Println("1. Add Member")
		fmt.Println("2. Add Book")
		fmt.Println("3. Remove Book")
		fmt.Println("4. List Members")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Borrow Book")
		fmt.Println("8. Return Book")
		fmt.Println("9. Exit")
		fmt.Println("==============================")
		fmt.Println("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("\033[H\033[2J") //clear the console
			c.AddMember(reader)
			time.Sleep(2 * time.Second)
		case "2":
			fmt.Print("\033[H\033[2J") //clear the console
			c.AddBook(reader)
			time.Sleep(2 * time.Second)
		case "3":
			fmt.Print("\033[H\033[2J") //clear the console
			c.RemoveBook(reader)
			time.Sleep(2 * time.Second)
		case "4":
			fmt.Print("\033[H\033[2J") //clear the console
			c.ListMembers()
			time.Sleep(2 * time.Second)
		case "5":
			fmt.Print("\033[H\033[2J") //clear the console
			c.ListAvailableBooks()
			time.Sleep(4 * time.Second)
		case "6":
			fmt.Print("\033[H\033[2J") //clear the console
			c.ListBorrowedBooks(reader)
			time.Sleep(4 * time.Second)
		case "7":
			fmt.Print("\033[H\033[2J") //clear the console
			c.BorrowBook(reader)
			time.Sleep(2 * time.Second)
		case "8":
			fmt.Print("\033[H\033[2J") //clear the console
			c.ReturnBook(reader)
			time.Sleep(2 * time.Second)
		case "9":
			fmt.Println("Exiting...")
			time.Sleep(2 * time.Second)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			time.Sleep(1 * time.Second)
		}
	}
}
