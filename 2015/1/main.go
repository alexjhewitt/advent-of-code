package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("levels.txt")
	if err != nil {
		panic(err)
	}
	floor := 0
	for _, c := range dat {
		fmt.Printf("%c", c)
	}
}
