package nks

import (
	"strconv"
	"time"
)

func getTickes() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
