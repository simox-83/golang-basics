package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func main() {
	log.Println("This is a log message") // very similar to the fmt.Println

	writeLog(INFO, "this is an INFO message")
	writeLog(WARNING, "this is a WARNING message")
	writeLog(ERROR, "this is an ERROR message")
	writeLog(FATAL, "this is a FATAL message, I am going to crash")
}

func writeLog(mt messageType, msg string) {

	//Open a file, create it if it doesn't exist, if it does append data. Pass the write only mode
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// set the log output to the file
	log.SetOutput(file)

	switch mt {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(msg)
	}

}
