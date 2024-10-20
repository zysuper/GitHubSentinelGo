package logger

import "fmt"

// Error 记录错误
func Error(args ...interface{}) {
    fmt.Println(args...)
}

// Info 记录信息
func Info(args ...interface{}) {
    fmt.Println(args...)
}
