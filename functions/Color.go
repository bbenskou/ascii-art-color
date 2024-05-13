package handler


var clr = struct {
	Reset   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	Gray    string
	White   string
}{
	Reset:   "\033[0m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m", // corrected from "\033[44m"
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	Gray:    "\033[37m",
	White:   "\033[97m",
}

func Color(str []string, color string) []string {
	cv := ""
	switch color {
	case "Red":
		cv = clr.Red
	case "Green":
		cv = clr.Green
	case "Yellow":
		cv = clr.Yellow
	case "Blue":
		cv = clr.Blue
	case "Magenta":
		cv = clr.Magenta
	case "Cyan":
		cv = clr.Cyan
	case "Gray":
		cv = clr.Gray
	case "White":
		cv = clr.White
	default:
		cv = clr.Reset
	}

	for i := 0; i < len(str); i++ {
		str[i] = cv + str[i] + clr.Reset
	}

	return str
}
