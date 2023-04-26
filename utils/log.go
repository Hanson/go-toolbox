package utils

import (
	"fmt"
	"github.com/hanson/gFile"
	"io"
	"log"
	"os"
	"time"
)

func KeepNewDateLogFile() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("err: %+v", err)
			}
		}()

		ticker := time.NewTicker(time.Minute)
		select {
		case <-ticker.C:
			now := time.Now()
			if now.Hour() == 0 && now.Minute() == 0 {
				log.SetOutput(GetMultiWriter())
			}
		}
	}()
}

func GetMultiWriter() io.Writer {
	err := gFile.CreateDirIfNotExists("logs", 0777)
	if err != nil {
		log.Printf("err: %+v", err)
		panic(err)
	}

	f, err := os.OpenFile(fmt.Sprintf("logs/%s.txt", time.Now().Format("2006-01-02")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return io.MultiWriter(os.Stdout, f)
}
