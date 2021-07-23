package main

import "fmt"

func main()  {
	forEach()
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