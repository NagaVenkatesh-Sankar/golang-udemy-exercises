package main

import (
	"encoding/json"
	// "errors"
	"fmt"
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
	// err := errors.New("errrror")
	if s, err := json.Marshal(users); err == nil {
		fmt.Println(string(s))
	}
}
