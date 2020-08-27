package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	//before sorting
	fmt.Println(xi)
	fmt.Println(xs)

	// integer sort
	sort.Ints(xi)
	// sort.IntSlice(xi).Sort()

	//integer reverse sort
	// sort.Sort(sort.Reverse(sort.IntSlice(xi)))

	//string sort
	sort.Strings(xs)

	//after sort
	fmt.Println(xi)
	fmt.Println(xs)
}
