package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/google/shlex"
)

type Library struct {
	books map[string]*Book 
}

type Book struct {
	isbn string
	title     string
	author    string
	year      string
	available bool
}

func (l *Library) addBook(parts []string) {
	if len(parts) < 5 {
		fmt.Println("USAGE: add {isbn} {title} {author} {year}")
		return
	}

	isbn := parts[1]
	title := parts[2]
	author := parts[3]
	year := parts[4]
	
	l.books[isbn] = &Book{isbn: isbn, title: title, author: author, year: year, available: true}
	fmt.Println("Book Added: ", l.books[isbn])
}

func (l *Library) updateBook(parts []string) {
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

	if book, ok := l.books[isbn]; ok {
		if title != "" {
			book.title = title
		}
		if author != "" {
			book.author = author
		}
		if year != "" {
			book.year = year
		}
		book.available = available
		fmt.Println("Book Updated:", book)
	} else {
		fmt.Println("Book not found for ISBN:", isbn)
	}
}

func (l *Library) deleteBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: delete {isbn}")
		return
	}
	isbn := parts[1]
	fmt.Println("Book delected: ", l.books[isbn])
	delete(l.books, isbn)
}

func (l Library) displayBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: display {isbn}")
		return
	}
	isbn := parts[1]
	if book, exists := l.books[isbn]; exists {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", book.isbn, book.title, book.author, book.year, book.available)
	}
}

func (l Library) listBooks() {
	fmt.Println("Displaying all books")
	for isbn, book := range l.books {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", isbn, book.title, book.author, book.year, book.available)
	}
}

func (l *Library) borrowBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: borrow {isbn}")
		return
	}

	isbn := parts[1]

	if book, ok := l.books[isbn]; ok {
		book.available = false
		fmt.Println("Book borrowed: ", book.title)
	} else {
		fmt.Println("Book does not exist")
	}
}

func (l *Library) returnBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: return {isbn}")
		return
	}

	isbn := parts[1]

	if book, ok := l.books[isbn]; ok {
		book.available = true
		fmt.Println("Book Returned: ", book.title)
	} else {
		fmt.Println("Book does not exist")
	}
}

func main() {
	lib := &Library{books: make(map[string]*Book)}
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
			lib.addBook(text)
		case "update":
			lib.updateBook(text)
		case "delete":
			lib.deleteBook(text)
		case "display":
			lib.displayBook(text)
		case "list":
			lib.listBooks()
		case "borrow":
			lib.borrowBook(text)
		case "return":
			lib.returnBook(text)
		case "quit":
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Unknown Command")
		}
	}
}