package helper

import "time"

func GetTime() (res int64) {
	res = time.Now().UTC().Unix()
	return
}
