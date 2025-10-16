package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/google/shlex"
)

type Book struct {
	isbn string
	title     string
	author    string
	year      string
	available bool
}

var books = make(map[string]*Book)

func addBook(parts []string) {
	if len(parts) < 5 {
		fmt.Println("USAGE: add {isbn} {title} {author} {year}")
		return
	}

	isbn := parts[1]
	title := parts[2]
	author := parts[3]
	year := parts[4]
	
	books[isbn] = &Book{isbn: isbn, title: title, author: author, year: year, available: true}
	fmt.Println("Book Added: ", books[isbn])
}

func updateBook(parts []string) {
	if len(parts) < 6 {
		fmt.Println("USAGE: update {isbn} {title} {author} {year} {available}")
		return
	}
	isbn := parts[1]
	title := parts[2]
	author := parts[3]
	year := parts[4]
	available, err := strconv.ParseBool(parts[5])

	if err != nil {
		fmt.Println("Invalid availability value (use true/false)")
		return
	}

	if book, ok := books[isbn]; ok {
		if title != "" {
			book.title = title
		}
		if author != "" {
			book.author = author
		}
		if year != "" {
			book.year = year
		}
		if year != "" {
			book.available = available
		}
	}
	
	//books[isbn] = &Book{isbn: isbn, title: title, author: author, year: year, available: available}
	fmt.Println("Book Updated: ", books[isbn])
}

func deleteBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: delete {isbn}")
		return
	}
	isbn := parts[1]
	fmt.Println("Book delected: ", books[isbn])
	delete(books, isbn)
}

func displayBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: display {isbn}")
		return
	}
	isbn := parts[1]
	if book, exists := books[isbn]; exists {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", book.isbn, book.title, book.author, book.year, book.available)
	}
}

func listBooks() {
	fmt.Println("Displaying all books")
	for isbn, book := range books {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", isbn, book.title, book.author, book.year, book.available)
	}
}

func borrowBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: borrow {isbn}")
		return
	}

	isbn := parts[1]

	if book, ok := books[isbn]; ok {
		book.available = false
		fmt.Println("Book borrowed: ", book.title)
	} else {
		fmt.Println("Book does not exist")
	}
}

func returnBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: return {isbn}")
		return
	}

	isbn := parts[1]

	if book, ok := books[isbn]; ok {
		book.available = true
		fmt.Println("Book Returned: ", book.title)
	} else {
		fmt.Println("Book does not exist")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("Enter Command: ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		text, err := shlex.Split(line)
		if err != nil {
			fmt.Println("parse error:", err)
			continue
		}
		cmd := text[0]

		switch cmd {
		case "add":
			addBook(text)
		case "update":
			updateBook(text)
		case "delete":
			deleteBook(text)
		case "display":
			displayBook(text)
		case "list":
			listBooks()
		case "borrow":
			borrowBook(text)
		case "return":
			returnBook(text)
		case "quit":
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Unknown Command")
		}
	}
}