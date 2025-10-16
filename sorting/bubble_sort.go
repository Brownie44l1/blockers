package main

import "fmt"

func sortBubble(arr []int) {
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
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	fmt.Println("let's test this shit out")
	arr := []int{-5, 3, 2, 1, -3, -3, 7, 2, -2}
	sortBubble(arr)

	arr1 := []int{64, 25, 12, 22, 11}
	selectionSort(arr1)
	fmt.Println(arr1)
}