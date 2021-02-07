package main

import (
	"fmt"

	"github.com/jdowning27/hyperchess/common"
)

func main() {
	fmt.Println("Welcome to HyperChess!")
	board := common.NewBoard()
	for _, row := range board {
		fmt.Println(row)
	}
}
