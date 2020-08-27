package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   int
}

func main() {

	u1 := user{First: "James", Age: 32}
	u2 := user{First: "Moneypenny", Age: 27}
	u3 := user{"M", 54}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	// create a file and the 'Write' method is used by the Writer interface
	// f, _ := os.Create("test.txt")

	// use the standard output - Print
	f := os.Stdout
	err := json.NewEncoder(f).Encode(users)
	if err != nil {
		fmt.Println("json encode error")
	}

}
