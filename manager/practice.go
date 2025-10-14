package main

import (
	"fmt"
	"reflect"
)

func variableAndTypes() {
	myInt := 5
	myFloat := 2.0
	myString := "Belmont"
	myBool := true
	//myComplex := 33
	myArray := [3]int{1, 2, 3}
	mySlice := []float64{1.1, 2.5, 1.5, 4.1}
	type myStruct struct {
		name string
		age int
	}
	p1 := myStruct {"belmont", 22}
	myPerson := struct {
		name string
		age  int
	}{
		name: "Alice",
		age:  30,
	}
	salaries := map[string]float64{
        "Alice": 50000.0,
        "Bob":   60000.0,
    }
	myPtr := &myInt 
	ch := make(chan int)
	/* var myInterface fmt.Stringer 
    myInterface = "here" */

	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(myFloat))
	fmt.Println(reflect.TypeOf(myString))
	fmt.Println(reflect.TypeOf(myBool))
	fmt.Println(reflect.TypeOf(myArray))
	fmt.Println(reflect.TypeOf(mySlice))
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(myPerson))
	fmt.Println(reflect.TypeOf(salaries))
	fmt.Println(reflect.TypeOf(myPtr))
	fmt.Println(reflect.TypeOf(ch))
	//fmt.Println(reflect.TypeOf(myInterface))
}

func swap() {
	a := 1
	b := 2
	fmt.Printf("Before Swap: a = %d, b = %d ", a, b)
	a, b = b, a
	fmt.Printf("After Swap: a = %d, b = %d ", a, b)
}

func revString() {
	a := "hello"
	fmt.Println("Before reverse: ", a)
	runes := []rune(a)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("After reverse: ", string(runes))
}

func fib(n int) {
	a, b := 0, 1
	if n >= 1 {
		fmt.Println(a)
	}
	if n >= 2 {
		fmt.Println(b)
	}
	
	for i := 2; i < n; i++ {
		next := a + b
		fmt.Println(next)
		a = b
		b = next
	}
}

func factorial(n int) {
	if n == 0 {
		fmt.Println("0")
	}
	if n == 1 {
		fmt.Println("1")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i 
	}
	fmt.Println(result)
}



/* func main() {
	variableAndTypes()
	swap()
	revString()
	fib(4)
	factorial(5)
} */