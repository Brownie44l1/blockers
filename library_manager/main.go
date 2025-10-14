package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Book struct {
	title     string
	author    string
	year      string
	available bool
}

var (
	books []Book
	bkMap = make(map[string]Book)
)

func addBook(parts []string) {
	if len(parts) < 4 {
		fmt.Println("USAGE: add {title} {author} {year}")
		return
	}
	title := parts[1]
	author := parts[2]
	year := parts[3]
	available := true

	//slice implementation
	b := Book{title: title, author: author, year: year, available: available}
	books = append(books, b)

	//map implementation
	bkMap[title] = b 
}

func updateBook(parts []string) {
	if len(parts) < 5 {
		fmt.Println("USAGE: update {title} {author} {year} {available}")
		return
	}
	title := parts[1]
	author := parts[2]
	year := parts[3]
	available, _ := strconv.ParseBool(parts[4])

	//map implementation
	bkMap[title] = Book{title: title, author: author, year: year, available: available}

	//slice implementation
	for i := range books {
		if books[i].title == title {
			books[i] = bkMap[title]
		}
	}
}

func deleteBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: delete {title}")
		return
	}
	title := parts[1]

	//map implementation
	delete(bkMap, title)

	//slice implementation
	for i, s := range books {
		if s.title == title {
			books = append(books[:i], books[i+1:]...)
		}
	}
}

func displayBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: display {title}")
		return
	}

	title := parts[1]
	
	//slice implementation
	for _, s := range books {
		if s.title == title {
			fmt.Printf("Title: %s, Author: %s, Year: %s, Available: %t", s.title, s.author, s.year, s.available)
		}
	}

	//map implementation
	if b, exist := bkMap[title]; exist {
		fmt.Printf("Title: %s, Author: %s, Year: %s, Available: %t", b.title, b.author, b.year, b.available)
	} else {
		fmt.Println("Book not found")
	}
}

func listBooks(parts []string) {
	if len(parts) < 1 {
		fmt.Println("USAGE: list")
		return
	}

	//slice implementation
	for _, s := range books {
		fmt.Printf("Title: %s, Author: %s, Year: %s, Available: %t", s.title, s.author, s.year, s.available)
	}

	//map implementation
	for i, b := range bkMap {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}", i, b.title, b.author, b.year, b.available)
	}
}

func borrowBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: borrow {title}")
		return
	}

	title := parts[1]

	//slice implementation (using index becuase we need to modify the value itself and not just a copy if we use the value/element.)
	for i := range books {
		if books[i].title == title {
			books[i].available = false
		}
	}

	//map implementation
	bk := bkMap[title]
	bk.available = false
	bkMap[title] = bk 
}

func returnBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: return {title}")
		return
	}

	title := parts[1]

	//slice implementation
	for i := range books {
		if books[i].title == title {
			books[i].available = true
		}
	}

	//map implementation
	bk := bkMap[title]
	bk.available = true
	bkMap[title] = bk 
}

func main() {
	scanner := bufio.NewScanner()
	
	for {
		if !scanner.Scan() {
			fmt.Println("Invalid input")
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		text := strings.Fields(line)
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
			listBooks(text)
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

/* addBook {title} {author} {year} — adds a new book and sets available = true.
updateBook {title} {new_author} {new_year} — updates details
deleteBook {title} — removes from collection.
displayBook {title} — prints book details.
listBooks — prints all stored books.
borrowBook {title} — marks book as unavailable.
returnBook {title} — marks as available again.
quit — exits. */