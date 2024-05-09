package handler

import "strings"

func NewLineCheck(str string) bool {
	cont := 0
	result := strings.ReplaceAll(str, "\\n", "\n")
	for i := 0; i < len(result); i++ {
		if result[i] == '\n' {
			cont++
		}
	}
	if cont == len(result) {
		return true
	}
	return false
}
