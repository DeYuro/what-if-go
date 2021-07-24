package main

import "fmt"

func main()  {
	m := map[int]int{1:1,2:2,3:3,4:4,5:5}
	fmt.Printf("Original: map %v with pointer %p \n", m, &m)
	byValue(m) //*map[int]int does not support indexing
	fmt.Printf("After by value func: map %v pointer %p \n", m, &m)

	//forEach()
}
// byValue func show what if map passed by value into func
func byValue(m map[int]int)  {
	fmt.Printf("Into by value func before change: array %v pointer %p \n", m, &m)
	m[2] = 10
	fmt.Printf("Into by value func after change: array %v pointer %p \n", m, &m)
}


// forEach show what happen when iterate over the map
// SPOILER: random order
func forEach()  {
	m := map[int]int{
		1: 10,
		2: 9,
		3: 8,
		4: 7,
		5: 6,
		6: 5,
		7: 4,
		8: 3,
		9: 2,
		10: 1,
	}

	fmt.Printf("map %v\n", m)
	for i,v := range m {
		fmt.Printf("index %d, value %d\n", i ,v )
	}
}