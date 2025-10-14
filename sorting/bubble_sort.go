package main

import "fmt"

func sort_bubble(arr []int) {
	n := len(arr)
	flag := true
	for flag {
		flag = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				flag = true
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		n -= 1 
	}
	fmt.Println(arr)
}

func main() {
	fmt.Println("let's test this shit out")
	arr := []int{-5, 3, 2, 1, -3, -3, 7, 2, -2}
	sort_bubble(arr)
}