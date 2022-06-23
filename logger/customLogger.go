package logger

import "log"

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func LogMessageInGreen(message string) {
	log.Println(green + message + reset)
}

func LogMessageInYellow(message string) {
	log.Println(yellow + message + reset)
}

func LogMessageInRed(message string) {
	log.Println(red + message + reset)
}

func LogFatalMessageInRed(message string, err error) {
	log.Fatalf("%s%s%v%s", red, message, err, reset)
}
