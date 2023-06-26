package utils

import (
	"strings"
	"time"
)

func GetTime() string {
	return strings.Split(time.Now().String(), ".")[0]
}
