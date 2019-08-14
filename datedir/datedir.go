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

func GetDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006/01/02")
}

func MkdirDate(src string) (string, error) {

	des := filepath.Join(src, getDay())
	return des, os.MkdirAll(des, 0755)
}
