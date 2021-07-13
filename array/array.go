package main

import "fmt"

func main()  {
	arr := [5]int{1,2,3,4,5}
	fmt.Printf("Original: array %v with pointer %p \n", arr, &arr)
	byValue(arr)
	fmt.Printf("After by value func: array %v pointer %p \n", arr, &arr)
	byPointer(&arr)
	fmt.Printf("After by pointer func: array %v pointer %p \n", arr, &arr)

	outOfRange(arr, 5)
}

// byValue func show what if array passed by value into func
func byValue(arr [5]int)  {
	fmt.Printf("Into by value func before change: array %v pointer %p \n", arr, &arr)
	arr[2] = 10
	fmt.Printf("Into by value func after change: array %v pointer %p \n", arr, &arr)
}

// byPointer func show what if array passed by pointer into func
func byPointer(arr *[5]int)  {
	fmt.Printf("Into by pointer func before change: array %v pointer %p \n", arr, &arr)
	arr[2] = 10
	fmt.Printf("Into by pointer func after change: array %v pointer %p \n", arr, &arr)
}

func outOfRange(arr [5]int, idx int)  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("try to get el from %v with len %d by idx %d: got %v \n", arr, len(arr), idx, r)
			fmt.Println("Recursive call with decremented idx")
			outOfRange(arr, idx - 1)
		} else {
			fmt.Printf("Goes ok when maxIdx < array length")
		}
	}()

	for i := 0; i <= idx; i++ {
		_ = arr[i]
	}
}