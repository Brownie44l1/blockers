package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name string
	age int
	location string
}

var students = make(map[string]Student)

func add() {
	if len(os.Args) < 5 {
		fmt.Println("USAGE <go run main.go add {name} {age} {location}")
	}
	name := os.Args[2]
	age, _ := strconv.Atoi(os.Args[3])
	location := os.Args[4]

	students[name] = Student{name: name, age: age, location: location}
	fmt.Println("Student added:", students[name])
}

func update() {
	if len(os.Args) < 5 {
		fmt.Println("USAGE <go run main.go update {name} [new data]")
	}
}

func displayStudent() {
	if len(os.Args) < 5 {
		fmt.Println("USAGE <go run main.go displayStudent {name}")
	}
}

func delete() {
	if len(os.Args) < 5 {
		fmt.Println("USAGE <go run main.go delete {name} ")
	}
}

func quit() {
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("USAGE <go run main.go [function]")
	}
	fmt.Println("Choose one of the following option to perform a command.")
	fmt.Println("1. add \n2. update \n3. displayStudent \n4. delete \n5. quit")

	for {
		switch os.Args[1] {
		case "add":
			add()
		case "update":
			update()
		case "displayStudent":
			displayStudent()
		case "delete":
			delete()
		case "quit":
			quit()
		}
	}
}