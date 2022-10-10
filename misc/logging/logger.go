package logging

import (
	"fmt"
	"io"
	"log"
	"mangrove/compiler/analysis/token"
	"os"
	"strconv"
	"time"
)

var Logger *log.Logger

func InitializeLogger() {
	filename := fmt.Sprintf("logs/%d-log.txt", time.Now().Unix())

	file, fileErr := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if fileErr != nil {
		panic(fileErr)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	Logger = log.New(multiWriter, "Mangrove Debug: ", log.LstdFlags)
}

func FatalUnexpectedTokenError(err int32, tok token.Token, message string, v ...any) {
	Logger.Fatalf("Unexpected Token '"+tok.Literal+"' on Line "+strconv.FormatInt(int64(tok.Line), 10)+" Column "+strconv.FormatInt(int64(tok.Line), 10)+"\n"+
		message,
		v)
}
