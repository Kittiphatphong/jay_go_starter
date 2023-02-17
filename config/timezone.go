package config

import (
	"go_starter/logs"
	"time"
	_ "time/tzdata"
)

func init() {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logs.Error(err)
		return
	}
	time.Local = location
}
