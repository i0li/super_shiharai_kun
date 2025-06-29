package config

import "os"

func initTimezone() {
	os.Setenv("TZ", "Asia/Tokyo")
}
