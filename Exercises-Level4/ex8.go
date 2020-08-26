package main

import (
	"fmt"
)

func main() {
	//create a map with slice values
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	}

	// add a slice to the map
	m["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	//loop through all map range and the slice values
	for k, v := range m {
		fmt.Println("Key : ", k)
		for i, s := range v {
			fmt.Printf("slice value at %d : %v\n", i, s)
		}
	}

	if _, ok := m["bond_james"]; ok {
		fmt.Println("Key found. Delete.")
		delete(m, "bond_james")
	}
	fmt.Println(m)
}
