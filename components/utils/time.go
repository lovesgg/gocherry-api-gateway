package utils

import "time"

/**
获取当前时间
 */
func GetNowTimeFormat() string {
	now := time.Now().Unix()
	formatTime := time.Unix(now, 0).Format("2006-01-02 15:04:05")
	return formatTime
}
