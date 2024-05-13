package handler

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GeneratePattern(arg string, cont string) {
	SplitFile := strings.Split(cont, "\n\n")
	SplitArg := strings.Split(arg, "\\n")
	result := make([][]string, len(SplitArg))
	for mak := 0; mak < len(SplitArg); mak++ {
		result[mak] = make([]string, 8)
	}
	for i := 0; i < len(SplitArg); i++ {
		for l := 0; l < 8; l++ {
			for j := 0; j < len(SplitArg[i]); j++ {
				if strings.ContainsAny(string(SplitArg[i][j]), os.Args[2]) {
					SplitCharacters := strings.Split(string(SplitFile[SplitArg[i][j]-32]), "\n")
					col := regexp.MustCompile(`^--color=(Red|Blue|Yellow|Green|Magenta|Cyan|Gray)$`)
					test := col.FindString(os.Args[1])
					col1 := regexp.MustCompile(`(Red|Blue|Yellow|Green|Magenta|Cyan|Gray)`)
					test1 := col1.FindString(test)
					if !col.MatchString(os.Args[1]) {
						fmt.Println("--color=<color>")
						os.Exit(1)
					}
					Color(SplitCharacters, test1)
					result[i][l] += SplitCharacters[l]
				} else {
					SplitCharacters := strings.Split(string(SplitFile[SplitArg[i][j]-32]), "\n")
					if SplitCharacters[l] == "" {
						result[i][l] += "      "
					}
					result[i][l] += SplitCharacters[l]
				}
			}
			if result[i][l] != "" {
				fmt.Println(string(result[i][l]))
			} else {
				fmt.Println()
				break
			}

		}
	}
}
