package main

import "fmt"

func main()  {
	emptySlice()
	emptySliceWithCapacity()
}
// emptySlice show what inside empty slice which created via var
func emptySlice()  {
	var s []int
	fmt.Printf("Empty slice created via `var`")
	sliceDescribe(s)
}
// emptySlice show what inside empty slice which created via var
func emptySliceWithCapacity()  {
	s := make([]int, 10)
	fmt.Printf("Empty slice created via `make`")
	sliceDescribe(s)
}

// sliceDescribe print slice content, len, cap
func sliceDescribe(s []int)  {
	fmt.Printf("content %v, len %d, cap %d\n", s, len(s), cap(s))
}