package main

import (
	"bufio"
	"fmt"
	"os"
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

	//slice implementation
	b := Book{title: title, author: author, year: year, available: true}
	books = append(books, b)

	//map implementation
	bkMap[title] = b 
	fmt.Println("Book added: ", b)
}

func updateBook(parts []string) {
	if len(parts) < 5 {
		fmt.Println("USAGE: update {title} {author} {year} {available}")
		return
	}
	title := parts[1]
	author := parts[2]
	year := parts[3]
	available, err := strconv.ParseBool(parts[4])

	if err != nil {
		fmt.Println("Invalid availability value (use true/false)")
		return
	}

	//map implementation
	bkMap[title] = Book{title: title, author: author, year: year, available: available}

	//slice implementation
	for i := range books {
		if books[i].title == title {
			books[i] = bkMap[title]
			break
		}
	}
	fmt.Println("Book Updated: ", bkMap[title])
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
			break
		}
	}
	fmt.Println("Book delected: ", bkMap[title])
}

func displayBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: display {title}")
		return
	}

	title := parts[1]
	
	//slice implementation
	fmt.Println("Displaying book requested")
	for _, s := range books {
		if s.title == title {
			fmt.Printf("Title: %s, Author: %s, Year: %s, Available: %t\n", s.title, s.author, s.year, s.available)
		}
	}
}

func listBooks() {
	fmt.Println("Displaying all books")
	//map implementation
	for i, b := range bkMap {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", i, b.title, b.author, b.year, b.available)
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
	if bk, ok := bkMap[title]; ok {
		bk.available = false
		bkMap[title] = bk
	}
	fmt.Println("Book borrowed: ", bkMap[title].title)
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
	if bk, ok := bkMap[title]; ok {
		bk.available = true
		bkMap[title] = bk
	}
	fmt.Println("Book Returned: ", bkMap[title].title)
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