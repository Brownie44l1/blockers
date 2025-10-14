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

func addStudent() {
	if len(os.Args) < 5 {
		fmt.Println("USAGE <go run main.go add {name} {age} {location}")
	}
	name := os.Args[2]
	age, _ := strconv.Atoi(os.Args[3])
	location := os.Args[4]

	studentStruct := Student{name: name, age: age, location: location}
	students[name] = studentStruct
	fmt.Println("Student added: ", students[name])
}

func updateStudent() {
	if len(os.Args) < 5 {
		fmt.Println("USAGE <go run main.go update {name} [new data]")
	}
	name := os.Args[2]
	age, _ := strconv.Atoi(os.Args[3])
	location := os.Args[4]
	students[name] = Student{name: name, age: age, location: location}
}

func displayStudent() {
	if len(os.Args) < 3 {
		fmt.Println("USAGE <go run main.go displayStudent {name}")
	}
	name := os.Args[2]
	if student, ok := students[name]; ok {
		fmt.Printf("Name: %s, Age: %d, Location: %s\n", student.name, student.age, student.location)
	} else {
		fmt.Println("Student not found")
	}
}

func deleteStudent() {
	if len(os.Args) < 3 {
		fmt.Println("USAGE <go run main.go delete {name} ")
	}
	name := os.Args[2]
	delete(students, name)
	fmt.Println("Deleted:", name)
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
			addStudent()
		case "update":
			updateStudent()
		case "displayStudent":
			displayStudent()
		case "delete":
			deleteStudent()
		case "quit":
			quit()
		default:
			fmt.Println("Unknown Command")
		}
	}
}