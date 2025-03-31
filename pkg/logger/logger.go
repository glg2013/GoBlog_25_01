// Package logger 日志相关
package logger

import "log"

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
