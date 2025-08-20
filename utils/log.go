package utils

import (
	"fmt"
	"github.com/hanson/gFile"
	"io"
	"log"
	"os"
	"path/filepath"
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
				log.SetOutput(GetMultiWriter(DAY))
			}
		}
	}()
}

const (
	DAY = iota
	HOUR
	MINUTE
)

func GetMultiWriter(split int) io.Writer {
	err := gFile.CreateDirIfNotExists("logs", 0777)
	if err != nil {
		log.Printf("err: %+v", err)
		panic(err)
	}

	var fileName string
	switch split {
	case DAY:
		fileName = fmt.Sprintf("logs/%s.txt", time.Now().Format("20060102"))
	case HOUR:
		fileName = fmt.Sprintf("logs/%s.txt", time.Now().Format("2006010215"))
	case MINUTE:
		fileName = fmt.Sprintf("logs/%s.txt", time.Now().Format("200601021504"))
	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return io.MultiWriter(os.Stdout, f)
}

func RemoveExpireLog(hours int) {
	dirPath := "logs"
	// 计算截止时间
	cutoffTime := time.Now().Add(-time.Duration(hours) * time.Hour)

	// 读取目录
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件的详细信息
		info, err := file.Info()
		if err != nil {
			continue
		}

		filePath := filepath.Join(dirPath, file.Name())

		// 检查文件修改时间
		if info.ModTime().Before(cutoffTime) {
			err := os.Remove(filePath)
			if err != nil {
				return
			}
		}
	}
}
