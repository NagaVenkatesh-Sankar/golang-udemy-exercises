package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// bs, _ := json.Marshal(p1)
	// include the error
	bs, _ := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))

}
