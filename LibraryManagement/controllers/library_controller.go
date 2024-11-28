package controllers

import (
	"LibraryManagement/models"
	"LibraryManagement/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var library = services.NewLibrary() 

func prompt() string {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			fmt.Println("Please enter a valid input. Try again.")
			continue
		}
		return input
	 }
 }

func addBook() {
	fmt.Print("Enter book ID: ")
    id, _ := strconv.Atoi(prompt())
    fmt.Print("Enter book title: ")
    title := prompt()
    fmt.Print("Enter book author: ")
    author := prompt()

    book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
    library.AddBook(book)
    fmt.Println("Book added successfully!")
	
}

func removeBook() {
	fmt.Print("Enter book ID: ")
    id, _ := strconv.Atoi(prompt())

	err := library.RemoveBook(id)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Book removed successfully!")
    }
}

func borrowBook() {
    fmt.Print("Enter book ID: ")
    bookID, _ := strconv.Atoi(prompt())
    fmt.Print("Enter member ID: ")
    memberID, _ := strconv.Atoi(prompt())

    err := library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Book borrowed successfully!")
    }
}

func returnBook() {
    fmt.Print("Enter book ID: ")
    bookID, _ := strconv.Atoi(prompt())
    fmt.Print("Enter member ID: ")
    memberID, _ := strconv.Atoi(prompt())

    err := library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Book returned successfully!")
    }
}

func listAvailableBooks() {
    books := library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
    } else {
        fmt.Println("Available Books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}

func listBorrowedBooks() {
    fmt.Print("Enter member ID: ")
    memberID, _ := strconv.Atoi(prompt())

    books := library.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println("No borrowed books.")
    } else {
        fmt.Println("Borrowed Books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}

func OperateServices() {
	for {
		fmt.Println("Library Management System")
        fmt.Println("1. Add Book")
        fmt.Println("2. Remove Book")
        fmt.Println("3. Borrow Book")
        fmt.Println("4. Return Book")
        fmt.Println("5. List Available Books")
        fmt.Println("6. List Borrowed Books by Member")
        fmt.Println("7. Exit")
        fmt.Print("Choose an option: ")


		choice , _ := strconv.Atoi(prompt())

		switch choice {
        case 1:
            addBook()
        case 2:
            removeBook()
        case 3:
            borrowBook()
        case 4:
            returnBook()
        case 5:
            listAvailableBooks()
        case 6:
            listBorrowedBooks()
        case 7:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
		
	}
}