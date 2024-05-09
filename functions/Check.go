package handler

import (
	"fmt"
	"os"
)

func Check(str string) {
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			fmt.Println("PROGRAM : Invalid characters")
			os.Exit(0)
		}
	}
}
