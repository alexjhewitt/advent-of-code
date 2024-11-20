package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("levels.txt")
	if err != nil {
		panic(err)
	}
	floor := 0
	for _, b := range dat {
		c := fmt.Sprintf("%c", b)
		fmt.Println(c)
		if c == "(" {
			floor += 1
		} else if c == ")" {
			floor -= 1
		} else {
			err = errors.New("invalid character detected")
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Printf("Floor level: %d\n", floor)
}
