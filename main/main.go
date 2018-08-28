package main

import (
	"fmt"
	"time"

	lumberjack "local/lumberjack"
)

var ljacklogger = &lumberjack.Logger{
	Filename:   "/media/jeromeku/easystore/workspace/go/src/local/lumberjack/main/test.log", //Sample file to upload
	MaxSize:    1,
	MaxBackups: 1,
	MaxAge:     1,
	Bucket:     "jeromeku_test_bucket", //GCS bucket
}

func main() {
	//ljacklogger.Write([]byte("Hello World!"))
	rotateEvery := 30
	fmt.Printf("Rotating every %d seconds...\n", rotateEvery)

	for {
		ljacklogger.Write([]byte("Hello World\n"))
		ljacklogger.Rotate()
		time.Sleep(time.Duration(rotateEvery) * time.Second)
	}
}
