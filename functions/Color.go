package handler


func Color(str []string,color string) []string {
var Reset = "\033[0m" 
var Red = "\033[31m" 
//var Green = "\033[32m" 
//var Yellow = "\033[33m" 
//var Blue = "\033[44m" 
//var Magenta = "\033[35m" 
//var Cyan = "\033[36m" 
//var Gray = "\033[37m" 
//var White = "\033[97m"
	for i :=0; i < len(str); i++ {
		str[i] = Red + str[i] + Reset
	}
	
	return str
}