package main

import (
	"fmt"
	"time"
)

func Substr(source interface{}, start_index int, length int) string {
	var s string
	if v, ok := source.(string); ok {
		s = v
	} else {
		s = fmt.Sprintf("%d", source)
	}
	if len([]rune(s)) > length {
		return string([]rune(s)[start_index : start_index+length])
	}
	return string([]rune(s)[start_index : start_index+len([]rune(s))])
}

func TimestampToDate(date_format string, timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format(date_format)
}

func DateToTimestamp(date_format string, date string) int {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(date_format, date, loc)
	return int(theTime.Unix())

}
