package main

import "fmt"

func main()  {
	empty()
	emptyWithCapacity()
	getElemInEmpty()
	getElemInEmptyCapacity()
	appendToEmpty()
	appendToEmptyWithCapacity()
}
// emptySlice show what inside empty slice which created via var
func empty()  {
	var s []int
	fmt.Printf("Empty slice created via `var`")
	describe(s)
}
// emptySlice show what inside empty slice which created via var
func emptyWithCapacity()  {
	s := make([]int, 10)
	fmt.Printf("Empty slice created via `make`")
	describe(s)
}

// describe print slice content, len, cap
func describe(s []int)  {
	fmt.Printf("content %v, len %d, cap %d\n", s, len(s), cap(s))
}

// getElemInEmpty show what if get el by index from slice created via `var`
func getElemInEmpty()  {
	var s []int

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("While getting el from slice created via `var` got panic: %v\n", r)
		}
	}()
	fmt.Println("\nGet elem from slice created via `var`")

	fmt.Println(s[2])
}

// getElemInEmptyCapacity show what if get el by index from slice created via `make`
func getElemInEmptyCapacity()  {
	s := make([]int, 10)
	fmt.Println("\nGet elem from slice created via `make`")
	fmt.Println(s[2])
}

func appendToEmpty()  {
	var s []int
	fmt.Println("\nAppend to empty slice created via `var`")
	describe(append(s, 42))
}

func appendToEmptyWithCapacity()  {
	s := make([]int, 10)
	fmt.Println("\nAppend to empty slice created via `make`")
	fmt.Println("slice before append")
	describe(s)
	fmt.Println("slice after append")
	describe(append(s, 42))
}