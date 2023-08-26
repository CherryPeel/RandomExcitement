package main

import (
	"fmt"
	"github.com/cherrypeel/pkg/web"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	logFileName := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logWriter := io.MultiWriter(logFile, os.Stdout)
	gin.DefaultWriter = logWriter

	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				deleteOldLogs()
			}
		}
	}()

	r.Static("./static/images", "./static/images")
	r.Static("./site", "./static/site")
	r.Static("./css", "./static/site/css")
	r.Static("./js", "./static/site/js")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("v1/allClassifies", web.AllClassify)
	r.Any("v1/refresh", web.Refresh)
	r.Any("v1/randomExcitement", web.RandomExcitement)
	r.Any("v1/random", web.RandomImgByReq)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})

	r.Run(":8080")

}

func deleteOldLogs() {
	logsDir := "logs"
	files, err := os.ReadDir(logsDir)
	if err != nil {
		logError("Error reading logs directory: %v", err)
		return
	}

	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour)
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}

		logFileTime, err := time.Parse("2006-01-02", fileInfo.Name()[:10])
		if err != nil {
			logError("Error parsing log file name: %v", err)
			continue
		}

		if logFileTime.Before(oneWeekAgo) {
			logFileName := filepath.Join(logsDir, fileInfo.Name())
			err := os.Remove(logFileName)
			if err != nil {
				logError("Error deleting old log file: %v", err)
			} else {
				logInfo("Deleted old log file: %s", logFileName)
			}
		}
	}
}

func logError(format string, a ...interface{}) {
	msg := fmt.Sprintf("[ERROR] "+format, a...)
	fmt.Println(msg)
	if gin.DefaultWriter != nil {
		log.New(gin.DefaultWriter, "", log.LstdFlags).Printf(msg)
	}
}

func logInfo(format string, a ...interface{}) {
	msg := fmt.Sprintf("[INFO] "+format, a...)
	fmt.Println(msg)
	if gin.DefaultWriter != nil {
		log.New(gin.DefaultWriter, "", log.LstdFlags).Printf(msg)
	}
}
