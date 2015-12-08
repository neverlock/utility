package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Printl(msg string, lfile string) error {
	if lfile == "" {
		log.Println(msg)
		return nil
	} else {
		file, err := os.OpenFile(lfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		t := time.Now()
		data := fmt.Sprintf("%s %s\n", t.Format("2006/01/02 15:04:05"), msg)
		_, err1 := file.Write([]byte(data))
		if err1 != nil {
			return err
		}
	}
	return nil
}
