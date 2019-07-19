package datedir

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_getDay(t *testing.T) {
	fmt.Println("Testing get Day", getDay())
}

func Test_MkdirDate(t *testing.T) {
	fmt.Println("Testing MkdirDate")
	ret := MkdirDate(".")
	if ret != nil {
		t.Errorf("Error occured: %v\n", ret)
	}
}

func Test_RemoveDir(t *testing.T) {
	var ret error
	fmt.Println("Clear testing DIR")
	currentTime := time.Now()
	c := currentTime.Format("2006/01/02")
	ct := strings.Split(c, "/")

	fmt.Println("Delete Dir :", c)
	ret = os.Remove(c)
	if ret != nil {
		t.Errorf("Error occured: %v\n", ret)
	}

	fmt.Printf("Delete Dir : %s/%s\n", ct[0], ct[1])
	ret = os.Remove(ct[0] + "/" + ct[1])
	if ret != nil {
		t.Errorf("Error occured: %v\n", ret)
	}

	fmt.Println("Delete Dir :", ct[0])
	ret = os.Remove(ct[0])
	if ret != nil {
		t.Errorf("Error occured: %v\n", ret)
	}

}
