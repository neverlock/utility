package datedir

import (
	"os"
	"path/filepath"
	"time"
)

func getDay() string {
	currentTime := time.Now()
	return currentTime.Format("2006/01/02")
}

func MkdirDate(src string) error {

	des := filepath.Join(src, getDay())
	return os.MkdirAll(des, 0755)
}
