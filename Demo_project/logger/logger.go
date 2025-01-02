package logger

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func Writelogs(){
	// Function to write logs to event.logs file.
	file, err := os.OpenFile("event.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil{
		log.Fatal("Unable to open file")
		return
	}	

	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}
// func main(){
// 	Writelogs()

// 	InfoLogger.Print("hi")
// }