package debug

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("log/debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
	}
	logger = log.New(file, "logger: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func Info(s ...interface{}) {
	logger.Println(s...)
}
