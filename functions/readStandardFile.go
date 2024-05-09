package handler

import (
	"fmt"
	"os"
)

func ReadStandardFile() string {
	cont, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("PROGRAM : ", err)
		os.Exit(0)
	}
	return string(cont)
}
