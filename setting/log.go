package setting

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func init() {
	os.MkdirAll("logs/GIN", os.ModePerm)
	os.MkdirAll("logs/Images", os.ModePerm)
	os.MkdirAll("logs/Socks5", os.ModePerm)
	os.MkdirAll("images", os.ModePerm)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录响应时间和请求路径
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		path := c.Request.URL.Path
		key := c.Request.URL.Query()
		clientIP := c.ClientIP()

		// 打印日志
		c.Writer.WriteHeaderNow()
		logMessage := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s?%s\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			latencyTime,
			clientIP,
			c.Request.Method,
			path,
			key,
		)

		go WriteLogFile(logMessage, "GIN", nil)
	}
}

func WriteFile(logMessage string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if !os.IsNotExist(err) {
		write := bufio.NewWriter(file)
		write.WriteString(logMessage)
		write.Flush()
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			log.Println("create file err:", err)
			return
		}
		write := bufio.NewWriter(file)
		write.WriteString(logMessage)
		write.Flush()
	}
}

func WriteLogFile(logs string, atype string, err error) {
	fileName := fmt.Sprintf("logs%c%s%clog-%s", os.PathSeparator, atype, os.PathSeparator, time.Now().Format("2006-01-02"))
	switch atype {
	case "GIN":
		WriteFile(logs, fileName)
	case "Images":
		if err != nil {
			logMessage := fmt.Sprintf("[yeyr2-newColly] %v | saving %s failed:%v\n", time.Now().Format("2006/01/02 - 15:04:05"), logs, err)
			WriteFile(logMessage, fileName)
			return
		}
		logMessage := fmt.Sprintf("[yeyr2-newColly] %v | saving %s success\n", time.Now().Format("2006/01/02 - 15:04:05"), logs)
		WriteFile(logMessage, fileName)
	case "Socks5":
		logs = fmt.Sprintf("[yeyr2-newColly] %v | %s\n", time.Now().Format("2006/01/02 - 15:04:05"), logs)
		WriteFile(logs, fileName)
	}

}
