package std

import "time"

func UnixTimeNowFn() int64 {
	return time.Now().Unix()
}
