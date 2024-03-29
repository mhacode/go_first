package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	// Sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice ", s)

	// Finding the index
	fmt.Println("Index vale :", strings.Index("GeeksforGeeks", "ks"))

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())

}
