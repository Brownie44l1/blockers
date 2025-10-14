package main

import "fmt"

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

func main() {

}

/* addBook {title} {author} {year} — adds a new book and sets available = true.
updateBook {title} {new_author} {new_year} — updates details
deleteBook {title} — removes from collection.
displayBook {title} — prints book details.
listBooks — prints all stored books.
borrowBook {title} — marks book as unavailable.
returnBook {title} — marks as available again.
quit — exits. */