package main

import (
	"fmt"
	"os"

	"handler/functions"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("PROGRAM : go run . <string>")
		return
	}
	arg := os.Args[3]
	handler.HandleInput(arg)
}
