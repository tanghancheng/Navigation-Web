package utils

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05.000")) // Format内即是你想转换的格式
	return []byte(stamp), nil
}
