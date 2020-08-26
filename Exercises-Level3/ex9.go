package main

import (
	"fmt"
)

func main() {
	favSport := "cricket"
	switch favSport {
	case "football":
		fmt.Println("Got a ball")
	case "hockey":
		fmt.Println("Got a hockey stick and ball")
	case "cricket":
		fmt.Println("Got a bat and ball")
	case "badminton":
		fmt.Println("Got a racquet and shuttle")
	}
}
