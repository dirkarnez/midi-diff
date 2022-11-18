package main

import (
	"fmt"

	"github.com/dirkarnez/gitdiff"
)

func main() {
	left, err := gitdiff.GetLeft()
	if err != nil {
		panic(err)
	}
	right, err := gitdiff.GetRight()
	if err != nil {
		panic(err)
	}

	fmt.Println(left, right)
}
