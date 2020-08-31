package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

// receiver function of 'sqrtError' that implements Error interface
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		// e := errors.New("Negative value error")
		e := fmt.Errorf("Negative value error - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
