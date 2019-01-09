package nks

import (
	"strconv"
	"time"
)

func getTicks() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
