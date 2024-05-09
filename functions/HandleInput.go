package handler

import (
	"fmt"
)

func HandleInput(arg string) {
	if NewLineCheck(arg) {
		for i := 0; i < len(arg)/2; i++ {
			fmt.Println()
		}
		return
	}
	Check(arg)
	cont := ReadStandardFile()
	GeneratePattern(arg, cont)
}
