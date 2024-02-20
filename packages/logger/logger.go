// 在 pkg/logger/logger.go
package logger

import (
	"fmt"
	"time"
)

// init函数在包被导入时自动执行，用于包的初始化。
func init() {
	fmt.Println("Logger package initialized")
}

// Log函数是这个日志库提供的一个示例功能。
func Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}
