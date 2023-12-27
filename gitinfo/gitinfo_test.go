package gitinfo

import (
	"fmt"
	"testing"
)

func Test_Info1(t *testing.T) {
	fmt.Println("Testing no file")
	msg, err := Info("./", "main")
	if err != nil {
		fmt.Println("Error: ", msg)
	} else {
		fmt.Println("Success: ", msg)
		t.Errorf("Error occured: %v , %s\n", err, msg)
	}
}
func Test_Info2(t *testing.T) {
	fmt.Println("Testing have file")
	msg, err := Info("../", "master")
	if err != nil {
		t.Errorf("Error occured: %v\n", err)
		fmt.Println("Error: ", msg)
	} else {
		fmt.Println("Success: ", msg)
	}
}
