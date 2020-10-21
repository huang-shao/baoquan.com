package util

import "time"

const TIME_FORMAT_FOUR  ="2006.01.02 15:04:05"

func TimeNow(format string)string  {
	return time.Now().Format(format)

}

func TimeFormat(sec int64,nsec int64,format string)string  {
	return time.Unix(sec,nsec).Format(format)

}