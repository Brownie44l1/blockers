package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string
	age int
	location string
}

var (
	students []Student 
	stMap = make(map[string]Student)
)

func addStudent(parts []string) {
	if len(parts) < 4 {
		fmt.Println("USAGE: add {name} {age} {location}")
		return
	}
	name := parts[1]
	age, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Age is invalid")
	}
	location := parts[3]

	s := Student{name: name, age: age, location: location}
	students = append(students, s)
	stMap[name] = s
	fmt.Println("Added: ", s)
}

func updateStudent(parts []string) {
	if len(parts) < 4 {
		fmt.Println("USAGE: update {name} {age} {location}")
		return
	}
	name := parts[1]
	age, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Age is invalid")
	}
	location := parts[3]
	stMap[name] = Student{name: name, age: age, location: location}
	for i := range students {
		if students[i].name == name {
			students[i] = stMap[name]
			break
		}
	}
	fmt.Println("Updated:", name)
}

func deleteStudent(parts []string) {
	if len(parts) < 3 {
		fmt.Println("USAGE: delete {name}")
		return
	}
	name := parts[2]
	delete(stMap, name)
	for i, s := range students {
		if s.name == name {
			students = append(students[:i], students[i+1:]...)
		}
	}
	fmt.Println("Deleted:", name)
}

func displayStudent(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: display {name|all}")
		return
	}
	if parts[1] == "all" {
		for _, s := range students {
			fmt.Printf("Name: %s, Age: %d, Location: %s\n", s.name, s.age, s.location)
		}
		return
	}
	name := parts[1]
	if s, ok := stMap[name]; ok {
		fmt.Printf("Name: %s, Age: %d, Location: %s\n", s.name, s.age, s.location)
	} else {
		fmt.Println("Student not found")
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
		parts := strings.Fields(line)
		cmd := parts[0]

		switch cmd {
		case "add":
			addStudent(parts)
		case "update":
			updateStudent(parts)
		case "displayStudent":
			displayStudent(parts)
		case "delete":
			deleteStudent(parts)
		case "quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown Command")
		}
	}
}