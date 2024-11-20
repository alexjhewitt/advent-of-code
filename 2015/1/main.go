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
	pos := 0
	for i, b := range dat {
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
		if floor == -1 {
			if pos != 0 {
				continue
			}
			pos = i
		}
	}
	fmt.Printf("Floor level: %d\n", floor)
	fmt.Println("Position where floor goes to -1: ", pos+1)
}
