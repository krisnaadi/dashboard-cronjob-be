package clocker

import (
	"time"
)

func Now() time.Time {

	// Menggunakan zona waktu GMT+8 (Asia/Singapore)
	loc, _ := time.LoadLocation("Asia/Makassar")
	t := time.Now().In(loc)
	return t
}

func Parse(oldTime time.Time) time.Time {

	// Menggunakan zona waktu GMT+8 (Asia/Singapore)
	loc, _ := time.LoadLocation("Asia/Makassar")
	newTime := oldTime.In(loc)
	return newTime
}
