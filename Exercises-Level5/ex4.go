package main

import (
	"fmt"
)

func main() {
	// anonymous struct
	s := struct {
		doors  int
		color  string
		luxury bool
		lights []string
	}{
		doors:  4,
		color:  "red",
		luxury: true,
		lights: []string{"front", "rear"},
	}

	fmt.Println(s)

	// s := struct {
	// 	vehicle struct {
	// 		doors int
	// 		color string
	// 	}
	// 	luxury bool
	// 	lights []string
	// }{
	// 	vehicle: vehicle{
	// 		doors: 4,
	// 		color: "red",
	// 	},
	// 	luxury: true,
	// 	lights: []string{"front", "rear"},
	// }
}
